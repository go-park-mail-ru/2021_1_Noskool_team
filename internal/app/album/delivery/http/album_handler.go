package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/pkg/response"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
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

	handler.router.HandleFunc("/{album_id:[0-9]+}", handler.GetAlbumByIDHandler).Methods("GET")

	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main of albums"))
	})

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

func (handler *AlbumsHandler) GetAlbumByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	albumID, _ := strconv.Atoi(mux.Vars(r)["album_id"])

	track, err := handler.albumsUsecase.GetAlbumByID(albumID)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByID: %v", err)
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
