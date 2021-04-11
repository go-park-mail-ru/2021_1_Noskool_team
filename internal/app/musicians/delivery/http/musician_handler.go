package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	oneDayTime = 86400
)

type MusiciansHandler struct {
	router          *mux.Router
	musicianUsecase musicians.Usecase
	logger          *logrus.Logger
	sessionsClient  client.AuthCheckerClient
}

func NewMusicHandler(r *mux.Router, config *configs.Config, usecase musicians.Usecase) *MusiciansHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &MusiciansHandler{
		router:          r,
		musicianUsecase: usecase,
		logger:          logrus.New(),
		sessionsClient:  client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	// /api/v1/musicians/
	middleware.ContentTypeJson(handler.router)
	handler.router.HandleFunc("/api/v1/musicians/bygenre/{genre}", handler.GetMusiciansByGenre)
	handler.router.HandleFunc("/api/v1/musicians/{musician_id:[0-9]+}", handler.GetMusicianByID)
	handler.router.HandleFunc("/api/v1/musicians/bytrack/{track_id:[0-9]+}", handler.GetMusicianByTrackID)
	handler.router.HandleFunc("/api/v1/musicians/byalbum/{album_id:[0-9]+}", handler.GetMusicianByAlbumID)
	handler.router.HandleFunc("/api/v1/musicians/byplaylist/{playlist_id:[0-9]+}", handler.GetMusicianByPlaylistID)

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

func (handler *MusiciansHandler) GetMusiciansByGenre(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	genre, ok := vars["genre"]
	if !ok {
		handler.logger.Errorf("Error get genre from query string")
		w.Write(response.FailedResponse(w, 500))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusiciansByGenre(genre)
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200)
}

func (handler *MusiciansHandler) GetMusicianByID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	musicianID, ok := vars["musician_id"]
	if !ok {
		handler.logger.Errorf("Error get musician_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musicianIDint, err := strconv.Atoi(musicianID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musician, err := handler.musicianUsecase.GetMusicianByID(musicianIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musician, 200)
}

func (handler *MusiciansHandler) GetMusicianByTrackID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	trackID, ok := vars["track_id"]
	if !ok {
		handler.logger.Errorf("Error get track_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	trackIDint, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusicianByTrackID(trackIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByTrackID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200)
}

func (handler *MusiciansHandler) GetMusicianByAlbumID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	albumID, ok := vars["album_id"]
	if !ok {
		handler.logger.Errorf("Error get album_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	albumIDint, err := strconv.Atoi(albumID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusicianByAlbumID(albumIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByAlbumID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200)
}

func (handler *MusiciansHandler) GetMusicianByPlaylistID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	playlistID, ok := vars["playlist_id"]
	if !ok {
		handler.logger.Errorf("Error get playlist_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	playlistIDint, err := strconv.Atoi(playlistID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusicianByPlaylistID(playlistIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByPlaylistID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200)
}
