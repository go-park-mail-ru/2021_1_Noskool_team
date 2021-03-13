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
	"os"
	"path/filepath"
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
	mediaFolder := fmt.Sprintf("%s", "./static")
	s.router.PathPrefix("/api/v1/data/").
		Handler(
			http.StripPrefix(
				"/api/v1/data/", http.FileServer(http.Dir(mediaFolder))))

	authMiddleware := middleware.NewSessionMiddleware(s.sessionsClient)
	cors := middleware.NewCORSMiddleware(s.config)
	s.router.Use(cors.CORS)
	s.router.HandleFunc("/api/v1/login",
		s.handleLogin()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/registrate",
		s.handleRegistrate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/logout",
		authMiddleware.CheckSessionMiddleware(s.handleLogout())).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/profile",
		authMiddleware.CheckSessionMiddleware(s.handleProfile())).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/profile/update",
		authMiddleware.CheckSessionMiddleware(s.handleUpdateProfile())).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/profile/avatar/upload",
		s.handleUpdateAvatar()).Methods(http.MethodPost, http.MethodOptions)

	s.router.Use(middleware.LoggingMiddleware)
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

func (s *ProfilesServer) HandleAuth(w http.ResponseWriter, r *http.Request) {
	SessionHash, _ := r.Cookie("session_id")
	_, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
	if err != nil {
		s.logger.Error("Пользователь не авторизован", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(200)
}

func (s *ProfilesServer) handleUpdateAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.logger.Info("handleUpdateAvatar")

		SessionHash, _ := r.Cookie("session_id")
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		if err != nil {
			s.logger.Error(err)
		}
		userIDfromCookie := session.ID
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)

		r.ParseMultipartForm(5 * 1024 * 1025)
		file, handler, err := r.FormFile("my_file")
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		defer file.Close()

		ext := filepath.Ext(handler.Filename)
		if ext == "" {
			fmt.Println("the file must have the extension")
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Загружаемый файл должен иметь расширение например img.phg"))
			return
		}
		path, _ := os.Getwd()
		photoPath := path + "/static/img/"
		newAvatarPath := photoPath + session.ID + ext
		fmt.Println(newAvatarPath)
		f, err := os.OpenFile(newAvatarPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере, не смогли создать файл картинки"))
			return
		}
		defer f.Close()
		io.Copy(f, file)
		s.db.User().UpdateAvatar(userIDfromCookieStr, "/api/v1/data/img/"+session.ID+ext)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleLogin() http.HandlerFunc {
	type request struct {
		Login    string `json:"nickname"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.logger.Info("starting handleLogin")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		u, err := s.db.User().FindByLogin(req.Login)

		if err != nil || !u.ComparePassword(req.Password) {
			fmt.Println(err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Некорректный email или пароль"))
			return
		}
		session, err := s.sessionsClient.Create(context.Background(), strconv.Itoa(u.ProfileID))
		fmt.Println("Result: = " + session.Status)
		if err != nil {
			s.logger.Errorf("Error in creating session: %v", err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Ошибка авторизации"))
			return
		}
		fmt.Println(session.ID)
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    session.Hash,
			Expires:  time.Now().Add(10000 * time.Hour),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleRegistrate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.logger.Info("starting handleRegistrate")
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
		}
		fmt.Println("req", req)
		u := &models.UserProfile{
			Email:    req.Email,
			Password: req.Password,
			Login:    req.Nickname,
			Avatar:   "/api/v1/data/img/default.png",
		}
		if err := s.db.User().Create(u); err != nil {
			var msgForUser string
			if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_email_key\"" {
				msgForUser = "Пользователь с таким email уже существует."
			}
			if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_nickname_key\"" {
				msgForUser = "Пользователь с таким nickname уже существует."
			}
			fmt.Println(err)
			s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf(msgForUser))
			return
		}
		fmt.Println("result of registration: ", u)
		u.Sanitize()
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *ProfilesServer) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		s.logger.Info("starting handleLogout")
		session, err := r.Cookie("session_id")
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		sessionID := session.Value
		result, err := s.sessionsClient.Delete(context.Background(), sessionID)
		fmt.Println("Result: = " + result.Status)
		if err != nil {
			s.logger.Errorf("Error in deleting session: %v", err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *ProfilesServer) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		SessionHash, _ := r.Cookie("session_id")
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		s.logger.Info("starting handleProfile")
		a, err := s.db.User().FindByID(session.ID)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		s.respond(w, r, http.StatusOK, a)
	}
}

func (s *ProfilesServer) handleUpdateProfile() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.logger.Info("starting handleUpdateProfile")

		SessionHash, _ := r.Cookie("session_id")
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		userIDfromCookie := session.ID
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)

		// fmt.Println(userIDfromURLstr)
		userForUpdates, err := s.db.User().FindByID(userIDfromCookieStr)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
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

		var msgForUser string
		if flagPassword {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				fmt.Println(err)
				if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_email_key\"" {
					msgForUser = "Пользователь с таким email уже существует."
				}
				if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_nickname_key\"" {
					msgForUser = "Пользователь с таким nickname уже существует."
				}
				s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf(msgForUser))
				return
			}
		} else {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				fmt.Println(err)
				if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_email_key\"" {
					msgForUser = "Пользователь с таким email уже существует."
				}
				if err.Error() == "pq: duplicate key value violates unique constraint \"profiles_nickname_key\"" {
					msgForUser = "Пользователь с таким nickname уже существует."
				}
				s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf(msgForUser))
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
		resp, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		w.Write(resp)
	}
}
