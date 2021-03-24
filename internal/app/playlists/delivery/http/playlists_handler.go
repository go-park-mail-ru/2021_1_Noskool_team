package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/playlists"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PlaylistsHandler struct {
	router         *mux.Router
	playlists      playlists.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewPlaylistsHandler(r *mux.Router, config *configs.Config, playlistsUsecase playlists.Usecase) *PlaylistsHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &PlaylistsHandler{
		router:         r,
		playlists:      playlistsUsecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}

	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}
	return handler
}

func ConfigLogger(handler *PlaylistsHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

