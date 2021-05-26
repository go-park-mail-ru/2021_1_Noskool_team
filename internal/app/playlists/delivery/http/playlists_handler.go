package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	albumModels "2021_1_Noskool_team/internal/app/album/models"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/playlists"
	playlistModels "2021_1_Noskool_team/internal/app/playlists/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	sessionModels "2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"2021_1_Noskool_team/internal/pkg/utility"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PlaylistsHandler struct {
	router           *mux.Router
	playlistsUsecase playlists.Usecase
	albumUsecae      album.Usecase
	logger           *logrus.Logger
	sessionsClient   client.AuthCheckerClient
}

func NewPlaylistsHandler(r *mux.Router, config *configs.Config, playlistsUsecase playlists.Usecase,
	albumUsecase album.Usecase) *PlaylistsHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &PlaylistsHandler{
		router:           r,
		playlistsUsecase: playlistsUsecase,
		albumUsecae:      albumUsecase,
		logger:           logrus.New(),
		sessionsClient:   client.NewSessionsClient(grpcCon),
	}

	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	authMiddlware := middleware.NewSessionMiddleware(handler.sessionsClient)

	handler.router.HandleFunc("/",
		authMiddlware.CheckSessionMiddleware(handler.CreatePlaylistHandler)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/",
		authMiddlware.CheckSessionMiddleware(handler.GetMediateka)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/top",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylists)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/top/notauth",
		handler.GetPlaylists).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.DeletePlaylistFromMediatekaHandler)).Methods(http.MethodDelete, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylistByIDHandler)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/getByUID/{playlist_uid}",
		handler.GetPlaylistByUIDHandler).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.AddPlaylistToMediatekaHandler)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/genre/{genre_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylistsByGenreID)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/picture",
		handler.UploadPlaylistPictureHandler).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/track/{track_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.AddTrackToPlaylist)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/track/{track_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.DeleteTrackFromPlaylist)).Methods(http.MethodDelete, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/description",
		authMiddlware.CheckSessionMiddleware(handler.UpdatePlaylistDescriptionHandler)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{playlist_id:[0-9]+}/title",
		authMiddlware.CheckSessionMiddleware(handler.UpdatePlaylistTittleHandler)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/user/{user_id:[0-9]+}",
		authMiddlware.CheckSessionMiddleware(handler.GetPlaylistsByUserID)).Methods(http.MethodGet, http.MethodOptions)

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

func (handler *PlaylistsHandler) GetPlaylists(w http.ResponseWriter, r *http.Request) {
	playlists, err := handler.playlistsUsecase.GetPlaylists()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, playlists, http.StatusOK, playlistModels.MarshalPlaylists)
}

func (handler *PlaylistsHandler) GetPlaylistsNotAuth(w http.ResponseWriter, r *http.Request) {
	playlists, err := handler.playlistsUsecase.GetPlaylists()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, playlists, http.StatusOK, playlistModels.MarshalPlaylists)
}

func (handler *PlaylistsHandler) CreatePlaylistHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	playlist := &playlistModels.Playlist{}
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
	response.SendCorrectResponse(w, playlist, http.StatusOK, playlistModels.MarshalPlaylist)
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

	for _, track := range playlist.Tracks {
		track.Albums = make([]*albumModels.Album, 0)
		albums, err := handler.albumUsecae.GetAlbumsByTrackID(track.TrackID)
		if err != nil {
			continue
		}
		track.Albums = append(track.Albums, &(*albums)[0])
	}

	response.SendCorrectResponse(w, playlist, http.StatusOK, playlistModels.MarshalPlaylist)
}

func (handler *PlaylistsHandler) GetPlaylistByUIDHandler(w http.ResponseWriter, r *http.Request) {
	playlistUID := mux.Vars(r)["playlist_uid"]

	playlist, err := handler.playlistsUsecase.GetPlaylistByUID(playlistUID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant find playlist with id = %s", playlistUID),
		})
		return
	}

	for _, track := range playlist.Tracks {
		track.Albums = make([]*albumModels.Album, 0)
		albums, err := handler.albumUsecae.GetAlbumsByTrackID(track.TrackID)
		if err != nil {
			continue
		}
		track.Albums = append(track.Albums, &(*albums)[0])
	}

	response.SendCorrectResponse(w, playlist, http.StatusOK, playlistModels.MarshalPlaylist)
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
	response.SendCorrectResponse(w, playlists, http.StatusOK, playlistModels.MarshalPlaylists)
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
	response.SendCorrectResponse(w, playlistsByGenreID, http.StatusOK, playlistModels.MarshalPlaylists)
}

func (handler *PlaylistsHandler) UploadPlaylistPictureHandler(w http.ResponseWriter, r *http.Request) {
	playlistID := mux.Vars(r)["playlist_id"]

	fileName, err := utility.SaveFile(r, "playlist_picture", "/static/img/playlists/", playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	fileNetPath := "/api/v1/music/data/img/playlists/" + *fileName
	playlistIDINT, err := strconv.Atoi(playlistID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	err = handler.playlistsUsecase.UploadPicture(playlistIDINT, fileNetPath)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *PlaylistsHandler) AddTrackToPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct playlist id",
		})
		return
	}
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct track id",
		})
		return
	}

	err = handler.playlistsUsecase.AddTrackToPlaylist(playlistID, trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant delete track %d from playlist  %d", trackID, playlistID),
		})
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *PlaylistsHandler) DeleteTrackFromPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct playlist id",
		})
		return
	}
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct track id",
		})
		return
	}

	err = handler.playlistsUsecase.DeleteTrackFromPlaylist(playlistID, trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNoContent,
			Message: fmt.Sprintf("Cant delete track %d from playlist  %d", trackID, playlistID),
		})
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *PlaylistsHandler) UpdatePlaylistTittleHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error(err)
		response.FailedResponse(w, http.StatusInternalServerError)
	}
	defer r.Body.Close()
	playlist := &playlistModels.Playlist{}
	err = playlist.UnmarshalJSON(body)
	if err != nil {
		handler.logger.Error(err)
		response.FailedResponse(w, http.StatusInternalServerError)
	}

	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct playlist id",
		})
		return
	}
	playlist.PlaylistID = playlistID
	userID, err := utility.CheckUserID(w, r, handler.logger)
	if err != nil {
		return
	}
	playlist.UserID = userID

	err = handler.playlistsUsecase.UpdatePlaylistTittle(playlist)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNotFound,
			Message: "Some error happened",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistsHandler) UpdatePlaylistDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.Error(err)
		response.FailedResponse(w, http.StatusInternalServerError)
	}
	defer r.Body.Close()
	playlist := &playlistModels.Playlist{}
	err = playlist.UnmarshalJSON(body)
	if err != nil {
		handler.logger.Error(err)
		response.FailedResponse(w, http.StatusInternalServerError)
	}

	playlistID, err := strconv.Atoi(mux.Vars(r)["playlist_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct playlist id",
		})
		return
	}
	playlist.PlaylistID = playlistID
	userID, err := utility.CheckUserID(w, r, handler.logger)
	if err != nil {
		return
	}
	playlist.UserID = userID

	err = handler.playlistsUsecase.UpdatePlaylistDescription(playlist)
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusNotFound,
			Message: "Some error happened",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistsHandler) GetPlaylistsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct user id",
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
	response.SendCorrectResponse(w, playlists, http.StatusOK, playlistModels.MarshalPlaylists)
}
