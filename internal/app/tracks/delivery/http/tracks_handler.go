package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/server"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type TracksHandler struct {
	router         *mux.Router
	tracksUsecase  tracks.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewTracksHandler(r *mux.Router, config *configs.Config, usecase tracks.Usecase) *TracksHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	handler := &TracksHandler{
		router:         r,
		tracksUsecase:  usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		fmt.Println(err)
	}
	handler.router.HandleFunc("/{track_id:[0-9]+}",
		handler.GetTrackByIdHandler).Methods("GET")
	handler.router.HandleFunc("/{track_tittle}", handler.GetTracksByTittle).Methods("GET")

	handler.router.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All tracks"))
	}).Methods("GET")
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main of tracks"))
	})

	return handler
}

func (handler *TracksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func ConfigLogger(handler *TracksHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

func (handler *TracksHandler) GetTrackByIdHandler(w http.ResponseWriter, r *http.Request) {
	trackID, _ := strconv.Atoi(mux.Vars(r)["track_id"])

	track, err := handler.tracksUsecase.GetTrackById(trackID)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		handler.logger.Errorf("Error in GetTrackById: %v", err)
		w.Write(server.FailedResponse())
		return
	}
	response, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(server.FailedResponse())
		return
	}
	w.Write(response)
}

func (handler *TracksHandler) GetTracksByTittle(w http.ResponseWriter, r *http.Request) {
	trackTittle := mux.Vars(r)["track_tittle"]

	track, err := handler.tracksUsecase.GetTracksByTittle(trackTittle)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		handler.logger.Errorf("Error in GetTracksByTittle: %v", err)
		w.Write(server.FailedResponse())
		return
	}
	response, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(server.FailedResponse())
		return
	}
	w.Write(response)
}
