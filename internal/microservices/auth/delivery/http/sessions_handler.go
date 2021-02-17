package http

import (
	"2021_1_Noskool_team/internal/microservices/auth"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
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

	userID := r.FormValue("user_id")

	session, err := handler.sessionsUsecase.CreateSession(userID)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   session.UserID,
		Expires: time.Now().Add(5 * time.Hour),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(session.UserID))
}

func (handler *SessionsHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")

	if err != nil {
		fmt.Println(err)
		return
	}

	//sessionID, _ := strconv.Atoi(session.Value)

	err = handler.sessionsUsecase.DeleteSession(session.Value)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("some error happened(("))
	} else {
		w.Write([]byte("cookie with id = " + session.Value + " was deleted"))

		session.Expires = time.Now().AddDate(0, 0, -5)
		http.SetCookie(w, session)
	}
}

func (handler *SessionsHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")

	if err != nil {
		fmt.Println(err)
	}

	//sessionID, _ := strconv.Atoi(session.Value)

	sess, err := handler.sessionsUsecase.CheckSession(session.Value)
	if err == nil && sess.UserID == session.Value {
		w.Write([]byte("Куку есть и id у нее = " + sess.UserID))
	} else {
		fmt.Println(err)
		w.Write([]byte("Куки нет"))
	}
}

func (handler *SessionsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
