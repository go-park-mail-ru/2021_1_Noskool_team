package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/playlists"
	"2021_1_Noskool_team/internal/app/playlists/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	sessionModels "2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strconv"
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

	authMiddlware := middleware.NewSessionMiddleware(handler.sessionsClient)

	handler.router.HandleFunc("/",
		authMiddlware.CheckSessionMiddleware(handler.CreatePlaylistHandler)).Methods(http.MethodPost)

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

func (handler *PlaylistsHandler) CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	playlist := &models.Playlist{}
	err = playlist.UnmarshalJSON(body)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	session, ok := r.Context().Value("user_id").(sessionModels.Result)
	if !ok {
		handler.logger.Error("Не получилось достать из конекста")
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct user id",
		})
		return
	}
	userID, err := strconv.Atoi(session.ID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error converting userID to int",
		})
		return
	}
	playlist.UserID = userID

	playlist, err = handler.playlists.CreatePlaylist(playlist)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Cant create playlist",
		})
		return
	}
	response.SendCorrectResponse(w, playlist, http.StatusOK)
}
