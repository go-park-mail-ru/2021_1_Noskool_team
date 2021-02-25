package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/music"
	"2021_1_Noskool_team/internal/app/music/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"encoding/json"
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
	musicUsecase   music.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}



func NewMusicHandler(config *configs.Config, usecase music.Usecase) *MusicHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	handler := &MusicHandler{
		router:         mux.NewRouter(),
		musicUsecase:   usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		fmt.Println(err)
	}

	authMiddleware := middleware.NewSessionMiddleware(handler.sessionsClient)

	checkAuth := handler.router.PathPrefix("/logged").Subrouter()
	checkAuth.Use(authMiddleware.CheckSessionMiddleware)
	checkAuth.HandleFunc("/getMusic", handler.GetMusic)
	checkAuth.HandleFunc("/deleteSession", handler.DeleteSession)

	handler.router.HandleFunc("/createSession", handler.CreateSession)
	handler.router.HandleFunc("/checkSession", handler.CheckSession)
	handler.router.HandleFunc("/getMusiciansByGenres", handler.GetMusiciansByGenres)
	handler.router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login page"))
	})
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main page"))
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
		w.Write([]byte("Кука есть и id у нее = " + strconv.Itoa(session.ID)))
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

func ConfigLogger(handler *MusicHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

func (handler *MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func (handler *MusicHandler) GetMusiciansByGenres(w http.ResponseWriter, r *http.Request) {
	genre := r.FormValue("genre")
	w.Header().Set("Content-Type", "application/json")
	musicians, err := handler.musicUsecase.GetMusiciansByGenres(genre)
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		w.Write(FailedResponse())
	}
	response, err := json.Marshal(musicians)
	if err != nil {
		handler.logger.Errorf("Error in marshalling json: %v", err)
		w.Write(FailedResponse())
	}
	w.Write(response)
}

func FailedResponse() []byte {
	response := models.FailedResponse{}
	response.ResultStatus = "failed"
	resp, _ := json.Marshal(response)
	return resp
}
