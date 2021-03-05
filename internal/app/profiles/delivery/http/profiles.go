package profiles

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"2021_1_Noskool_team/internal/app/profiles/repository"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	oneDayTime = 86400
)

// ProfilesServer ...
type ProfilesServer struct {
	config         *configs.Config
	logger         *logrus.Logger
	router         *mux.Router
	db             *repository.Store
	sessionsClient client.AuthCheckerClient
}

// New ...
func New(config *configs.Config) *ProfilesServer {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}
	return &ProfilesServer{
		config:         config,
		logger:         logrus.New(),
		router:         mux.NewRouter(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
}

// Start ...
func (s *ProfilesServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureDB(); err != nil {
		return err
	}

	s.logger.Info("starting profile server")
	return http.ListenAndServe(s.config.ProfilesServerAddr, s.router)
}

func (s *ProfilesServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *ProfilesServer) configureRouter() {

	authMiddleware := middleware.NewSessionMiddleware(s.sessionsClient)

	//TODO апишку надо бы сделать в формате /api/vi/prifile/login ну кароч в REST-стиле
	s.router.HandleFunc("/api/v1/login", s.handleLogin())
	s.router.HandleFunc("/api/v1/registrate", s.handleRegistrate()).Methods("POST")
	s.router.HandleFunc("/api/v1/logout", authMiddleware.CheckSessionMiddleware(s.handleLogout()))
	s.router.HandleFunc("/api/v1/profile", authMiddleware.CheckSessionMiddleware(s.handleProfile()))
	s.router.HandleFunc("/api/v1/profile/{user_id:[0-9]+}", authMiddleware.CheckSessionMiddleware(s.handleUpdateProfile()))

	CORSMiddleware := middleware.NewCORSMiddleware(s.config)
	s.router.Use(CORSMiddleware.CORS)
	s.router.Use(middleware.PanicMiddleware)
}

func (s *ProfilesServer) configureDB() error {
	st := repository.New(s.config)
	if err := st.Open(); err != nil {
		return err
	}
	s.db = st
	return nil
}

func (s *ProfilesServer) handleLogin() http.HandlerFunc {
	type request struct {
		Login    string `json:"nickname"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleLogin")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.db.User().FindByLogin(req.Login)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("incorrect email or password"))
			return
		}
		session, err := s.sessionsClient.Create(context.Background(), u.ProfileID)
		fmt.Println("Result: = " + session.Status)

		if err != nil {
			s.logger.Errorf("Error in creating session: %v", err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("authorization error"))
			return
		}

		cookie := &http.Cookie{
			Path:    "/",
			Name:    "session_id",
			Value:   strconv.Itoa(session.ID),
			Expires: time.Now().Add(oneDayTime * time.Second),
		}

		http.SetCookie(w, cookie)

		io.WriteString(w, "successfully login for user with id "+fmt.Sprint(session.ID))
	}
}

func (s *ProfilesServer) handleRegistrate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleRegistrate")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}
		fmt.Println("req", req)

		u := &models.UserProfile{
			Email:    req.Email,
			Password: req.Password,
			Login:    req.Nickname,
		}

		fmt.Println("u", u)
		if err := s.db.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
		io.WriteString(w, "registrate")
	}
}

func (s *ProfilesServer) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleLogout")

		session, err := r.Cookie("session_id")

		if err != nil {
			s.logger.Errorf("Error in parsing cookie: %v", err)
			return
		}

		sessionID, _ := strconv.Atoi(session.Value)

		result, err := s.sessionsClient.Delete(context.Background(), sessionID)
		fmt.Println("Result: = " + result.Status)

		if err != nil {
			s.logger.Errorf("Error in deleting session: %v", err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("some bad"))
		}
		cookie := &http.Cookie{
			Path:    "/",
			Name:    "session_id",
			Value:   "",
			Expires: time.Unix(0, 0),
		}
		http.SetCookie(w, cookie)
		io.WriteString(w, "cookie with id = "+session.Value+" was deleted")
	}
}

func (s *ProfilesServer) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Cookie("session_id")
		s.logger.Info("starting handleProfile")

		a, err := s.db.User().FindByID(userID.Value)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("user with id="+userID.Value+" not found"))
			return
		}
		credentialsFromDB, err := json.Marshal(a)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		io.WriteString(w, string(credentialsFromDB))
	}
}

func (s *ProfilesServer) handleUpdateProfile() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleUpdateProfile")

		userIDfromURL, _ := strconv.Atoi(mux.Vars(r)["user_id"])
		userIDfromURLstr := fmt.Sprint(userIDfromURL)
		fmt.Println(userIDfromURLstr)

		userForUpdates, err := s.db.User().FindByID(userIDfromURLstr)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("user with id="+userIDfromURLstr+" not found"))
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
		}

		flagPassword := false
		if req.Email != "" {
			userForUpdates.Email = req.Email
		}
		if req.Nickname != "" {
			userForUpdates.Login = req.Nickname
		}
		if req.Password != "" {
			userForUpdates.Password = req.Password
			flagPassword = true
		}

		fmt.Println(userForUpdates)

		if flagPassword {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		} else {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				s.error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
		}

		userForUpdates.Sanitize()

		s.respond(w, r, http.StatusCreated, userForUpdates)
	}
}

func (s *ProfilesServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *ProfilesServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		fmt.Println(err)
	}
}
