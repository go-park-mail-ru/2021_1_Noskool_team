package profiles

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"2021_1_Noskool_team/internal/app/profiles/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ProfilesServer ...
type ProfilesServer struct {
	config *configs.Config
	logger *logrus.Logger
	router *mux.Router
	db     *repository.Store
}

// New ...
func New(config *configs.Config) *ProfilesServer {
	return &ProfilesServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
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
	s.router.HandleFunc("/login", s.handleLogin())
	s.router.HandleFunc("/registrate", s.handleRegistrate()).Methods("POST")
	s.router.HandleFunc("/logout", s.handleLogout())
	s.router.HandleFunc("/profile", s.handleProfile())
}

func (s *ProfilesServer) configureDB() error {
	st := repository.New(s.config.ProfileDB)
	if err := st.Open(); err != nil {
		return err
	}
	s.db = st
	return nil
}

func (s *ProfilesServer) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleLogin")
		io.WriteString(w, "login")
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
		io.WriteString(w, "logout")
	}
}

func (s *ProfilesServer) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleProfile")
		io.WriteString(w, "profile")
	}
}

func (s *ProfilesServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *ProfilesServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
