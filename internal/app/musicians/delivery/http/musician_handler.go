package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

const (
	oneDayTime = 86400
)

type MusiciansHandler struct {
	router         *mux.Router
	musicUsecase   musicians.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewMusicHandler(r *mux.Router, config *configs.Config, usecase musicians.Usecase) *MusiciansHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &MusiciansHandler{
		router:         r,
		musicUsecase:   usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	handler.router.HandleFunc("/{genre}/", handler.GetMusiciansByGenres)
	handler.router.HandleFunc("/{musician_id:[0-9]+}", handler.GetMusicByIDHandler)
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main page of music"))
	})
	return handler
}

func ConfigLogger(handler *MusiciansHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

func (handler *MusiciansHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func (handler *MusiciansHandler) GetMusiciansByGenres(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	genre := mux.Vars(r)["genre"]
	musicians, err := handler.musicUsecase.GetMusiciansByGenres(genre)
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	resp, err := json.Marshal(musicians)
	if err != nil {
		handler.logger.Errorf("Error in marshalling json: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	w.Write(resp)
}

func (handler *MusiciansHandler) GetMusicByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	musicianID, _ := strconv.Atoi(mux.Vars(r)["musician_id"])

	track, err := handler.musicUsecase.GetMusicianByID(musicianID)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	resp, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	w.Write(resp)
}


