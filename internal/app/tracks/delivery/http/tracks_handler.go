package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/response"
	"encoding/json"
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
		logrus.Error(err)
	}

	handler := &TracksHandler{
		router:         r,
		tracksUsecase:  usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	handler.router.HandleFunc("/{track_id:[0-9]+}", handler.GetTrackByIDHandler)
	handler.router.HandleFunc("/{track_tittle}", handler.GetTracksByTittle).Methods("GET")
	handler.router.HandleFunc("/musician/{musician_id:[0-9]+}", handler.GetTrackByMusicianID).Methods("GET")
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

func (handler *TracksHandler) GetTrackByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 500))
		return
	}

	track, err := handler.tracksUsecase.GetTrackByID(trackID)
	if err != nil {
		handler.logger.Errorf("Error in GetTrackByID: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	resp, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	w.Write(resp)
}

func (handler *TracksHandler) GetTracksByTittle(w http.ResponseWriter, r *http.Request) {
	trackTittle := mux.Vars(r)["track_tittle"]

	w.Header().Set("Content-Type", "application/json")

	track, err := handler.tracksUsecase.GetTracksByTittle(trackTittle)
	if err != nil {
		handler.logger.Errorf("Error in GetTracksByTittle: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	resp, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	w.Write(resp)
}

func (handler *TracksHandler) GetTrackByMusicianID(w http.ResponseWriter, r *http.Request) {
	musicianID, err := strconv.Atoi(mux.Vars(r)["musician_id"])
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 500))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	track, err := handler.tracksUsecase.GetTrackByMusicianID(musicianID)
	if err != nil {
		handler.logger.Errorf("Error in GetTrackByMusicianID: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	res, err := json.Marshal(track)
	if err != nil {
		handler.logger.Errorf("Error in marshalling: %v", err)
		w.Write(response.FailedResponse(w, http.StatusInternalServerError))
		return
	}
	w.Write(res)
}
