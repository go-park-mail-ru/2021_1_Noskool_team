package http

import (
	"2021_1_Noskool_team/internal/app/music"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

type MusicHandler struct {
	router         *mux.Router
	musicUsecase   *music.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewMusicHandler(usecase music.Usecase) *MusicHandler {
	grpcCon, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	handler := &MusicHandler{
		router:         mux.NewRouter(),
		musicUsecase:   &usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler)
	if err != nil {
		fmt.Println(err)
	}

	handler.router.HandleFunc("/getMusic", handler.GetMusic)
	handler.router.HandleFunc("/createSession", handler.CreateSession)
	handler.router.HandleFunc("/deleteSession", handler.DeleteSession)
	handler.router.HandleFunc("/checkSession", handler.CheckSession)
	handler.router.HandleFunc("/getMusic", handler.GetMusic)
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	return handler
}

func (handler *MusicHandler) GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Music"))
}

func (handler *MusicHandler) CreateSession(w http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(r.FormValue("user_id"))

	session, err := handler.sessionsClient.Create(context.Background(), userID)
	fmt.Println("Result: = " + session.Status)
	if err != nil {
		handler.logger.Errorf("Error in creating session: %v", err)
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(session.ID),
		Expires: time.Now().Add(5 * time.Hour),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(strconv.Itoa(session.ID)))
}

func (handler *MusicHandler) CheckSession(w http.ResponseWriter, r *http.Request) {

	sessionID, err := r.Cookie("session_id")

	if err != nil {
		handler.logger.Errorf("Error in parsing cookie: %v", err)
		w.Write([]byte("Куки нет, нужно редиректнуть на авторизацию"))
		return
	}

	userID, _ := strconv.Atoi(sessionID.Value)
	session, err := handler.sessionsClient.Check(context.Background(), userID)
	fmt.Println("Result: = " + session.Status)

	if err == nil && session.ID == userID {
		w.Write([]byte("Куку есть и id у нее = " + strconv.Itoa(session.ID)))
	} else {
		handler.logger.Errorf("Error in checking session: %v", err)
		w.Write([]byte("Куки нет"))
	}
}

func (handler *MusicHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")

	if err != nil {
		handler.logger.Errorf("Error in parsing cookie: %v", err)
		return
	}

	sessionID, _ := strconv.Atoi(session.Value)

	result, err := handler.sessionsClient.Delete(context.Background(), sessionID)
	fmt.Println("Result: = " + result.Status)
	if err != nil {
		handler.logger.Errorf("Error in deleting session: %v", err)
		w.Write([]byte("some error happened(("))
	} else {
		w.Write([]byte("cookie with id = " + session.Value + " was deleted"))

		session.Expires = time.Now().AddDate(0, 0, -5)
		http.SetCookie(w, session)
	}
}

func ConfigLogger(handler *MusicHandler) error {
	//level, err := logrus.ParseLevel(logrus.DebugLevel)
	//if err != nil {
	//	return err
	//}

	handler.logger.SetLevel(logrus.DebugLevel)
	return nil
}

func (handler *MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
