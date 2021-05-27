package profiles

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/profiles"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	authModels "2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/monitoring"
	"2021_1_Noskool_team/internal/pkg/response"
	"2021_1_Noskool_team/internal/pkg/utility"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/microcosm-cc/bluemonday"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// ProfilesServer ...
type ProfilesServer struct {
	config         *configs.Config
	logger         *logrus.Logger
	router         *mux.Router
	sessionsClient client.AuthCheckerClient
	profUsecase    profiles.Usecase
	sanitizer      *bluemonday.Policy
}

// New ...
func New(config *configs.Config, profUsecase profiles.Usecase, sanitizer *bluemonday.Policy) *ProfilesServer {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}
	return &ProfilesServer{
		config:         config,
		logger:         logrus.New(),
		router:         mux.NewRouter(),
		sessionsClient: client.NewSessionsClient(grpcCon),
		profUsecase:    profUsecase,
		sanitizer:      sanitizer,
	}
}

// Start ...
func (s *ProfilesServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

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
	mediaFolder := "./static"
	s.router.PathPrefix("/api/v1/user/data/").
		Handler(
			http.StripPrefix(
				"/api/v1/user/data/", http.FileServer(http.Dir(mediaFolder))))

	metricks := monitoring.RegisterMetrics(s.router)

	s.router.Use(middleware.LoggingMiddleware(metricks))

	authMiddleware := middleware.NewSessionMiddleware(s.sessionsClient)
	cors := middleware.NewCORSMiddleware(s.config)
	s.router.Use(cors.CORS)
	s.router.HandleFunc("/api/v1/user/login",
		s.handleLogin()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/registrate",
		s.handleRegistrate()).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/logout",
		authMiddleware.CheckSessionMiddleware(s.handleLogout())).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/csrf",
		authMiddleware.CheckSessionMiddleware(s.CreateCSRFHandler)).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile",
		authMiddleware.CheckSessionMiddleware(middleware.CheckCSRFMiddleware(s.handleProfile()))).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/update",
		authMiddleware.CheckSessionMiddleware(s.handleUpdateProfile())).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/avatar/upload",
		authMiddleware.CheckSessionMiddleware(s.handleUpdateAvatar())).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/update/password",
		authMiddleware.CheckSessionMiddleware(s.handleUpdatePassword())).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/{other_user_id:[0-9]+}",
		authMiddleware.CheckSessionMiddleware(s.GetOtherUserPage)).Methods(http.MethodGet, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/{other_user_id:[0-9]+}/subscribe",
		authMiddleware.CheckSessionMiddleware(s.SubscribeMeToSomebody)).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/{other_user_id:[0-9]+}/unsubscribe",
		authMiddleware.CheckSessionMiddleware(s.UnSubscribeMeToSomebody)).Methods(http.MethodPost, http.MethodOptions)
	s.router.HandleFunc("/api/v1/user/profile/search",
		authMiddleware.CheckSessionMiddleware(s.SearchContent)).Methods(http.MethodGet, http.MethodOptions)

	s.router.Use(middleware.PanicMiddleware(metricks))
	s.router.Use(middleware.ContentTypeJson)
}

func (s *ProfilesServer) SearchContent(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("search")
	searchQuery = s.sanitizer.Sanitize(searchQuery)

	fmt.Println(searchQuery)
	otherUsers, _ := s.profUsecase.SearchTracks(searchQuery)
	response.SendCorrectResponse(w, otherUsers, http.StatusOK, models.MarshalOtherUsers)
}

func (s *ProfilesServer) GetOtherUserPage(w http.ResponseWriter, r *http.Request) {
	userID, err := utility.CheckUserID(w, r, s.logger)
	if err != nil {
		return
	}
	otherUserID, err := strconv.Atoi(mux.Vars(r)["other_user_id"])
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct other user id",
		})
		return
	}
	otherUser, err := s.profUsecase.GetOtherUserPage(userID, otherUserID)
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Нет такого юзера",
		})
		return
	}
	response.SendCorrectResponse(w, otherUser, http.StatusOK, models.MarshalOtherUserFullInformation)
}

func (s *ProfilesServer) SubscribeMeToSomebody(w http.ResponseWriter, r *http.Request) {
	userID, err := utility.CheckUserID(w, r, s.logger)
	if err != nil {
		return
	}
	otherUserID, err := strconv.Atoi(mux.Vars(r)["other_user_id"])
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct other user id",
		})
		return
	}
	err = s.profUsecase.SubscribeMeToSomebody(userID, otherUserID)
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct other user id",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ProfilesServer) UnSubscribeMeToSomebody(w http.ResponseWriter, r *http.Request) {
	userID, err := utility.CheckUserID(w, r, s.logger)
	if err != nil {
		return
	}
	otherUserID, err := strconv.Atoi(mux.Vars(r)["other_user_id"])
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct other user id",
		})
		return
	}
	err = s.profUsecase.UnsubscribeMeToSomebody(userID, otherUserID)
	if err != nil {
		s.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct other user id",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ProfilesServer) CreateCSRFHandler(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("user_id").(authModels.Result)
	if !ok {
		s.logger.Error("Не получилось достать из конекста")
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct user id",
		})
		return
	}
	b := session.ID + session.Hash
	csrfToken := utility.CreateCSRFToken(b)
	csrfCookie := &http.Cookie{
		Name:    "csrf",
		Value:   csrfToken,
		Path:    "/",
		Expires: time.Now().Add(30 * time.Minute),
	}
	http.SetCookie(w, csrfCookie)
	w.Header().Set("csrf", csrfToken)
}

func (s *ProfilesServer) HandleAuth(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("HandleAuth")
	SessionHash, err := r.Cookie("session_id")
	if err != nil {
		s.logger.Error(err)
		s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
	}
	_, err = s.sessionsClient.Check(context.Background(), SessionHash.Value)
	if err != nil {
		s.logger.Error("Пользователь не авторизован", err)
		s.respond(w, r, http.StatusUnauthorized, nil)
		return
	}
	s.respond(w, r, http.StatusOK, nil)
}

func (s *ProfilesServer) handleUpdateAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("handleUpdateAvatar")
		w.Header().Set("Content-Type", "application/json")
		SessionHash, err := r.Cookie("session_id")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		if err != nil {
			s.logger.Error("Пользователь не авторизован", err)
			s.error(w, r, http.StatusUnauthorized, nil)
			return
		}
		userIDfromCookie := session.ID
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)

		_ = r.ParseMultipartForm(5 * 1024 * 1025)
		file, handler, err := r.FormFile("my_file")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		defer file.Close()

		ext := filepath.Ext(handler.Filename)
		if ext == "" {
			s.logger.Error("the file must have the extension")
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Загружаемый файл должен иметь расширение например img.phg"))
			return
		}
		path, _ := os.Getwd()
		photoPath := path + "/static/img/"
		newAvatarPath := photoPath + session.ID + ext
		s.logger.Info("newAvatarPath: ", newAvatarPath)
		f, err := os.OpenFile(newAvatarPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере, не смогли создать файл картинки"))
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file)
		s.profUsecase.UpdateAvatar(userIDfromCookieStr, "/api/v1/user/data/img/"+session.ID+ext)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleLogin")
		w.Header().Set("Content-Type", "application/json")
		req := &models.RequestLogin{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		req.Sanitize(s.sanitizer)
		fmt.Println("req.Login", req.Login, "req.Password", req.Password)
		u, err := s.profUsecase.FindByLogin(req.Login)
		fmt.Println("user", u)
		if err != nil || !u.ComparePassword(req.Password) {
			s.logger.Error(err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Некорректный nickname или пароль"))
			return
		}
		session, err := s.sessionsClient.Create(context.Background(), strconv.Itoa(u.ProfileID))
		s.logger.Info("Result: = " + session.Status)
		if err != nil {
			s.logger.Errorf("Error in creating session: %v", err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Ошибка авторизации"))
			return
		}
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    session.Hash,
			Path:     "/",
			Expires:  time.Now().Add(10000 * time.Hour),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleRegistrate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleRegistrate")
		w.Header().Set("Content-Type", "application/json")

		// TODO: проверка авторизован ли уже пользователь???

		req := &models.ProfileForRegistrate{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
		}
		err = json.Unmarshal(body, &req)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
		}
		req.Sanitize(s.sanitizer)

		u := &models.UserProfile{
			Email:    req.Email,
			Password: req.Password,
			Login:    req.Nickname,
			Avatar:   "/api/v1/user/data/img/default.png",
		}
		if err := s.profUsecase.Create(u); err != nil {
			s.logger.Error(err)
			msg, httpCode := checkDBerr(err)
			s.error(w, r, httpCode, fmt.Errorf(msg))
			return
		}
		s.logger.Info("result of registration: ", u)
		u.Sanitize()
		session, err := s.sessionsClient.Create(context.Background(), strconv.Itoa(u.ProfileID))
		s.logger.Info("Result: = " + session.Status)
		if err != nil {
			s.logger.Errorf("Error in creating session: %v", err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Ошибка авторизации"))
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		s.logger.Info("starting handleLogout")
		session, err := r.Cookie("session_id")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		result, err := s.sessionsClient.Delete(context.Background(), session.Value)
		s.logger.Info("Result: = " + result.Status)
		if err != nil {
			s.logger.Errorf("Error in deleting session: %v", err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *ProfilesServer) handleProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleProfile")
		w.Header().Set("Content-Type", "application/json")
		SessionHash, err := r.Cookie("session_id")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
		}
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
		}
		profile, err := s.profUsecase.FindByID(session.ID)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		s.respond(w, r, http.StatusOK, profile)
	}
}

func (s *ProfilesServer) handleUpdateProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleUpdateProfile")
		w.Header().Set("Content-Type", "application/json")

		SessionHash, err := r.Cookie("session_id")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusUnauthorized, nil)
			return
		}
		userIDfromCookie := session.ID
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)

		userForUpdates, err := s.profUsecase.FindByID(userIDfromCookieStr)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		req := &models.ProfileRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
		}
		req.Sanitize(s.sanitizer)
		if req.Email != "" {
			userForUpdates.Email = req.Email
		}
		if req.Name != "" {
			userForUpdates.Name = req.Name
		}
		if req.Surname != "" {
			userForUpdates.Surname = req.Surname
		}
		if req.Nickname != "" {
			userForUpdates.Login = req.Nickname
		}
		if len(req.FavoriteGenre) != 0 {
			userForUpdates.FavoriteGenre = req.FavoriteGenre
		}

		s.logger.Info("userForUpdates: ", userForUpdates)

		if err := s.profUsecase.Update(userForUpdates); err != nil {
			msg, httpCode := checkDBerr(err)
			s.error(w, r, httpCode, fmt.Errorf(msg))
			return
		}

		userForUpdates.Sanitize()
		s.respond(w, r, http.StatusCreated, userForUpdates)
	}
}

func (s *ProfilesServer) handleUpdatePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("starting handleUpdatePassword")
		w.Header().Set("Content-Type", "application/json")

		SessionHash, err := r.Cookie("session_id")
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusInternalServerError, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		session, err := s.sessionsClient.Check(context.Background(), SessionHash.Value)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusUnauthorized, nil)
			return
		}
		userIDfromCookie := session.ID
		userIDfromCookieStr := fmt.Sprint(userIDfromCookie)

		userForUpdates, err := s.profUsecase.FindByID(userIDfromCookieStr)
		if err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Не удалось найти пользователя"))
			return
		}
		req := &models.ChangePassword{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.logger.Error(err)
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Cервер не смог обработать информацию :("))
			return
		}
		req.Sanitize(s.sanitizer)

		if req.OldPassword == req.NewPassword {
			s.error(w, r, http.StatusBadRequest, fmt.Errorf("Новый пароль совпадает со старым."))
			return
		}

		if err != nil || !userForUpdates.ComparePassword(req.OldPassword) {
			s.logger.Error(err)
			s.error(w, r, http.StatusUnauthorized, fmt.Errorf("Введен неверный старый пароль"))
			return
		}

		if err := s.profUsecase.UpdatePassword(userForUpdates.ProfileID, req.NewPassword); err != nil {
			fmt.Println("2>", err)
			//msg, httpCode := checkDBerr(err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		userForUpdates.Sanitize()
		s.respond(w, r, http.StatusOK, nil)
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
			s.logger.Error(err)
			s.error(w, r, http.StatusUnprocessableEntity, fmt.Errorf("Ошибка на сервере :("))
			return
		}
		_, _ = w.Write(resp)
	}
}

func checkDBerr(err error) (string, int) {
	fmt.Println("checkDBerr: ", err)
	var msgForUser string
	var httpCode int

	var rightJSON map[string]interface{}
	errorMsgInJSON := json.Unmarshal([]byte(err.Error()), &rightJSON)

	switch {
	case err == models.ErrConstraintViolationEmail:
		msgForUser = "Пользователь с таким email уже существует."
		httpCode = http.StatusUnprocessableEntity
	case err == models.ErrConstraintViolationNickname:
		msgForUser = "Пользователь с таким nickname уже существует."
		httpCode = http.StatusUnprocessableEntity
	case errorMsgInJSON == nil:
		msgForUser = err.Error()
		httpCode = http.StatusBadRequest
	default:
		msgForUser = "Неопознаная ошибка на севере, ухх..."
		httpCode = http.StatusInternalServerError
	}
	return msgForUser, httpCode
}
