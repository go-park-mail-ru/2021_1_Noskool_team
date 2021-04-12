package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type AlbumsHandler struct {
	router         *mux.Router
	albumsUsecase  album.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewAlbumsHandler(r *mux.Router, config *configs.Config, usecase album.Usecase) *AlbumsHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &AlbumsHandler{
		router:         r,
		albumsUsecase:  usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}

	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	middleware.ContentTypeJson(handler.router)
	handler.router.HandleFunc("/api/v1/album/{album_id:[0-9]+}", handler.GetAlbumByID).Methods("GET")
	handler.router.HandleFunc("/api/v1/album/bymusician/{musician_id:[0-9]+}", handler.GetAlbumsByMusicianID).Methods("GET")
	handler.router.HandleFunc("/api/v1/album/bytrack/{track_id:[0-9]+}", handler.GetAlbumsByTrackID).Methods("GET")

	return handler
}

func (handler *AlbumsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func ConfigLogger(handler *AlbumsHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	handler.logger.SetLevel(level)
	return nil
}

func (handler *AlbumsHandler) GetAlbumByID(w http.ResponseWriter, r *http.Request) {
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
	album, err := handler.albumsUsecase.GetAlbumByID(albumIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, album, 200)
}

func (handler *AlbumsHandler) GetAlbumsByMusicianID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	musicianID, ok := vars["musician_id"]
	if !ok {
		handler.logger.Errorf("Error get album_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	musicianIDint, err := strconv.Atoi(musicianID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	album, err := handler.albumsUsecase.GetAlbumsByMusicianID(musicianIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumsByMusicianID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, album, 200)
}

func (handler *AlbumsHandler) GetAlbumsByTrackID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	trackID, ok := vars["track_id"]
	if !ok {
		handler.logger.Errorf("Error get album_id from query string")
		w.Write(response.FailedResponse(w, 400))
		return
	}
	trackIDint, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		w.Write(response.FailedResponse(w, 400))
		return
	}
	album, err := handler.albumsUsecase.GetAlbumsByTrackID(trackIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByTrackID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, album, 200)
}
