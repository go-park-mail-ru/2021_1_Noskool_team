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
	//return http.ListenAndServeTLS(s.config.ProfilesServerAddr, "server.crt", "server.key", s.router)
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
	cors := middleware.NewCORSMiddleware(s.config)
	s.router.Use(cors.CORS)
	s.router.HandleFunc("/api/v1/picture", mainPage)
	s.router.HandleFunc("/api/v1/test", s.handleUpdateAvatar())
	s.router.HandleFunc("/api/v1/login", s.handleLogin()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/registrate", s.handleRegistrate())
	s.router.HandleFunc("/api/v1/logout", authMiddleware.CheckSessionMiddleware(s.handleLogout()))
	s.router.HandleFunc("/api/v1/profile", authMiddleware.CheckSessionMiddleware(s.handleProfile()))
	s.router.HandleFunc("/api/v1/profile/{user_id:[0-9]+}", authMiddleware.CheckSessionMiddleware(s.handleUpdateProfile()))
	s.router.HandleFunc("/api/v1/profile/avatar/{user_id:[0-9]+}", s.handleUpdateAvatar())

	s.router.PathPrefix("/api/v1/data/").
		Handler(
			http.StripPrefix(
				"/api/v1/data/", http.FileServer(http.Dir("./static"))))

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

// на фронте нужна такая форма
var uploadFormTmpl = []byte(`
<html>
        <body>
        <form action="/api/v1/test" method="post" enctype="multipart/form-data">
                Image: <input type="file" name="my_file">
                <input type="submit" value="Upload">
        </form>
        </body>
</html>
`)

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write(uploadFormTmpl)
}

func (s *ProfilesServer) handleUpdateAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		s.logger.Info("handleUpdateAvatar")
		userIDfromURL, _ := strconv.Atoi(mux.Vars(r)["user_id"])
		//userIDfromURL := 1
		userIDfromURLstr := fmt.Sprint(userIDfromURL)
		userIDfromCookie, _ := r.Cookie("session_id")
		//userIDfromCookie := 1
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)
		if userIDfromURLstr != userIDfromCookieStr {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка доступа"))
			return
		}
		r.ParseMultipartForm(5 * 1024 * 1025)
		file, handler, err := r.FormFile("my_file")
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		defer file.Close()
		// fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
		// fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
		ext := filepath.Ext(handler.Filename)
		if ext == "" {
			fmt.Println("the file must have the extension")
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Загружаемый файл должен иметь расширение например img.phg"))
			return
		}
		newAvatarPath := "/api/v1/data/img/" + userIDfromCookieStr + ext
		f, err := os.OpenFile(newAvatarPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		defer f.Close()
		io.Copy(f, file)
		s.db.User().UpdateAvatar(userIDfromCookieStr, newAvatarPath)
		s.respond(w, r, http.StatusOK, nil)
		// io.WriteString(w, "Успех")
	}
}

func (s *ProfilesServer) handleLogin() http.HandlerFunc {
	type request struct {
		Login    string `json:"nickname"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
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
		session, err := s.sessionsClient.Create(context.Background(), u.ProfileID)
		fmt.Println("Result: = " + session.Status)
		if err != nil {
			s.logger.Errorf("Error in creating session: %v", err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Ошибка авторизации"))
			return
		}
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    strconv.Itoa(session.ID),
			Expires:  time.Now().Add(10000 * time.Hour),
			Secure:   false,
			HttpOnly: false,
		}
		http.SetCookie(w, &cookie)
		// w.WriteHeader(http.StatusOK)
		s.respond(w, r, http.StatusOK, nil)
		//body, _ := json.Marshal(map[string]string{"message": "Successful login."})
		//w.Write(body)
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
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		fmt.Println("req", req)
		u := &models.UserProfile{
			Email:    req.Email,
			Password: req.Password,
			Login:    req.Nickname,
		}
		if err := s.db.User().Create(u); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		fmt.Println("result of registration: ", u)
		u.Sanitize()
		// resp, err := json.Marshal(u)
		// if err != nil {
		// 	fmt.Println(err)
		// 	s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf("Ошибка на сервере :("))
		// 	return
		// }
		s.respond(w, r, http.StatusOK, u)
		//body, _ := json.Marshal(map[string]string{"message": "Successful login."})
		//io.WriteString(w, "registrate")
		//w.WriteHeader(http.StatusOK)
		//w.Write(body)
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
		sessionID, _ := strconv.Atoi(session.Value)
		result, err := s.sessionsClient.Delete(context.Background(), sessionID)
		fmt.Println("Result: = " + result.Status)
		if err != nil {
			s.logger.Errorf("Error in deleting session: %v", err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		//cookie := &http.Cookie{
		//      Path:     "/",
		//      Name:     "session_id",
		//      Value:    "",
		//      Expires:  time.Unix(0, 0),
		//      HttpOnly: true,
		//      Secure:   false,
		//}
		//http.SetCookie(w, cookie)
		//w.WriteHeader(http.StatusOK)
		//w.Write([]byte("{\"status\": \"OK\"}"))
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		w.WriteHeader(http.StatusOK)
	}
}

func (s *ProfilesServer) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Cookie("session_id")
		s.logger.Info("starting handleProfile")
		a, err := s.db.User().FindByID(userID.Value)
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
		userIDfromURL, _ := strconv.Atoi(mux.Vars(r)["user_id"])
		userIDfromURLstr := fmt.Sprint(userIDfromURL)
		userIDfromCookie, _ := r.Cookie("session_id")
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie.Value)
		// fmt.Println(">>>>>>>>>>", userIDfromURLstr, userIDfromCookieStr)
		if userIDfromURLstr != userIDfromCookieStr {
			fmt.Println("user_id from the cookie and from the url do not match")
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка доступа"))
			return
		}
		// fmt.Println(userIDfromURLstr)
		userForUpdates, err := s.db.User().FindByID(userIDfromURLstr)
		if err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			fmt.Println(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
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
		if flagPassword {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				fmt.Println(err)
				s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
				return
			}
		} else {
			if err := s.db.User().Update(userForUpdates, flagPassword); err != nil {
				fmt.Println(err)
				s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
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
