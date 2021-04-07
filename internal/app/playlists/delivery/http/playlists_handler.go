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
	"2021_1_Noskool_team/internal/pkg/utility"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PlaylistsHandler struct {
	router           *mux.Router
	playlistsUsecase playlists.Usecase
	logger           *logrus.Logger
	sessionsClient   client.AuthCheckerClient
}

func NewPlaylistsHandler(r *mux.Router, config *configs.Config, playlistsUsecase playlists.Usecase) *PlaylistsHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &PlaylistsHandler{
		router:           r,
		playlistsUsecase: playlistsUsecase,
		logger:           logrus.New(),
		sessionsClient:   client.NewSessionsClient(grpcCon),
	}

	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	authMiddlware := middleware.NewSessionMiddleware(handler.sessionsClient)

	handler.router.HandleFunc("/",
		authMiddlware.CheckSessionMiddleware(handler.CreatePlaylistHandler)).Methods(http.MethodPost)
	handler.router.HandleFunc("/",
		authMiddlware.CheckSessionMiddleware(handler.GetMediateka)).Methods(http.MethodGet)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.DeletePlaylistFromMediatekaHandler)).Methods(http.MethodDelete)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylistByIDHandler)).Methods(http.MethodGet)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.AddPlaylistToMediatekaHandler)).Methods(http.MethodPost)
	handler.router.HandleFunc("/genre/{genre_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylistsByGenreID)).Methods(http.MethodGet)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/picture",
		handler.UploadPlaylistPictureHandler).Methods(http.MethodPost)

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
	playlist, err = handler.playlistsUsecase.CreatePlaylist(playlist)
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

func (handler *PlaylistsHandler) DeletePlaylistFromMediatekaHandler(w http.ResponseWriter, r *http.Request) {
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
	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	err = handler.playlistsUsecase.DeletePlaylistFromUser(userID, playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error while deleting playlist",
		})
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *PlaylistsHandler) GetPlaylistByIDHandler(w http.ResponseWriter, r *http.Request) {
	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	playlist, err := handler.playlistsUsecase.GetPlaylistByID(playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant find playlist with id = %d", playlistID),
		})
		return
	}
	response.SendCorrectResponse(w, playlist, http.StatusOK)
}

func (handler *PlaylistsHandler) AddPlaylistToMediatekaHandler(w http.ResponseWriter, r *http.Request) {
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
	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	err = handler.playlistsUsecase.AddPlaylistToMediateka(userID, playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant add playlist with id = %d to mediateka", playlistID),
		})
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *PlaylistsHandler) GetMediateka(w http.ResponseWriter, r *http.Request) {
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
	playlists, err := handler.playlistsUsecase.GetMediateka(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant get mediateka for user with id = %d", userID),
		})
		return
	}
	response.SendCorrectResponse(w, playlists, http.StatusOK)
}

func (handler *PlaylistsHandler) GetPlaylistsByGenreID(w http.ResponseWriter, r *http.Request) {
	genreID, err := strconv.Atoi(mux.Vars(r)["genre_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct genre id",
		})
		return
	}
	playlistsByGenreID, err := handler.playlistsUsecase.GetPlaylistsByGenreID(genreID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant get playlists by genre id = %d", genreID),
		})
		return
	}
	response.SendCorrectResponse(w, playlistsByGenreID, http.StatusOK)
}

func (handler *PlaylistsHandler) UploadPlaylistPictureHandler(w http.ResponseWriter, r *http.Request) {
	playlistID := mux.Vars(r)["playlist_id"]

	fileName, err := utility.SaveFile(r, "playlist_picture", "/static/img/playlists/", playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	fileNetPath := "/api/v1/data/img/playlists/" + *fileName
	playlistIDINT, err := strconv.Atoi(playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	err = handler.playlistsUsecase.UploadAudio(playlistIDINT, fileNetPath)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	response.SendEmptyBody(w, http.StatusOK)
}
