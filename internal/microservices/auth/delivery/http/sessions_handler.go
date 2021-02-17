package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"testWorkWithAuth/internal/microservices/auth"
	"time"
)

type SessionsHandler struct {
	router          *mux.Router
	sessionsUsecase auth.Usecase
	logger          *logrus.Logger
}

func NewSessionHandler(usecase auth.Usecase) *SessionsHandler {
	handler := &SessionsHandler{
		router:          mux.NewRouter(),
		sessionsUsecase: usecase,
		logger:          logrus.New(),
	}

	handler.router.HandleFunc("/create_session", handler.CreateSession)
	handler.router.HandleFunc("/delete_session", handler.DeleteSession)
	handler.router.HandleFunc("/check_session", handler.CheckSession)
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("sessions"))
	})
	return handler
}

func (handler *SessionsHandler) CreateSession(w http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(r.FormValue("user_id"))

	session, err := handler.sessionsUsecase.CreateSession(userID)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Name: "session_id",
		Value: strconv.Itoa(session.UserID),
		Expires: time.Now().Add(5 * time.Hour),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(strconv.Itoa(session.UserID)))
}

func (handler *SessionsHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteSession"))
}

func (handler *SessionsHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CheckSession"))
}

func (handler *SessionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
