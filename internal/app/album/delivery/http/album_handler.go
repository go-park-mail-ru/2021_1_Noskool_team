package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	albumModels "2021_1_Noskool_team/internal/app/album/models"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/tracks"
	models0 "2021_1_Noskool_team/internal/app/tracks/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"2021_1_Noskool_team/internal/pkg/utility"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type AlbumsHandler struct {
	router         *mux.Router
	albumsUsecase  album.Usecase
	tracksUsecase  tracks.Usecase
	musUsecase     musicians.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewAlbumsHandler(r *mux.Router, config *configs.Config, usecase album.Usecase,
	tracksUsecase tracks.Usecase, musUsecase musicians.Usecase) *AlbumsHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &AlbumsHandler{
		router:         r,
		albumsUsecase:  usecase,
		tracksUsecase:  tracksUsecase,
		musUsecase:     musUsecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}

	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	authmiddlware := middleware.NewSessionMiddleware(handler.sessionsClient)
	middleware.ContentTypeJson(handler.router)
	handler.router.HandleFunc("/favorites",
		authmiddlware.CheckSessionMiddleware(middleware.CheckCSRFMiddleware(handler.GetFavoriteAlbums))).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{album_id:[0-9]+}", handler.GetAlbumByID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bymusician/{musician_id:[0-9]+}", handler.GetAlbumsByMusicianID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bytrack/{track_id:[0-9]+}", handler.GetAlbumsByTrackID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/{album_id:[0-9]+}/favorites",
		authmiddlware.CheckSessionMiddleware(handler.AddDeleteAlbumToFavorites)).Methods("POST", http.MethodOptions)
	handler.router.HandleFunc("/{album_id:[0-9]+}/mediateka",
		authmiddlware.CheckSessionMiddleware(
			middleware.CheckCSRFMiddleware(handler.AddDeleteAlbumToFavorites))).Methods("POST", http.MethodOptions)

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
	albumWithTracks := ConvertAlumToFullAlbum(album)
	albumWithTracks.Tracks, _ = handler.tracksUsecase.GetTracksByAlbumID(albumWithTracks.AlbumID)
	albumWithTracks.Musician, _ = handler.musUsecase.GetMusicianByAlbumID(albumWithTracks.AlbumID)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByID: %v", err)
		w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, albumWithTracks, 200, MarshalAlbumWithExtraInform)
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
	response.SendCorrectResponse(w, album, 200, albumModels.MarshalAlbums)
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
	response.SendCorrectResponse(w, album, 200, albumModels.MarshalAlbums)
}

func (handler *AlbumsHandler) AddDeleteAlbumToMediateka(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("user_id").(models.Result)
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
	trackID, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	addOrDelete := r.URL.Query().Get("type")
	if addOrDelete == "add" {
		err = handler.albumsUsecase.AddAlbumToMediateka(userID, trackID)
	} else if addOrDelete == "delete" {
		err = handler.albumsUsecase.DeleteAlbumFromMediateka(userID, trackID)
	}
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *AlbumsHandler) AddDeleteAlbumToFavorites(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("user_id").(models.Result)
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
	trackID, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	addOrDelete := r.URL.Query().Get("type")
	if addOrDelete == "add" {
		err = handler.albumsUsecase.AddAlbumToFavorites(userID, trackID)
	} else if addOrDelete == "delete" {
		err = handler.albumsUsecase.DelteAlbumFromFavorites(userID, trackID)
	}
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *AlbumsHandler) GetFavoriteAlbums(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("user_id").(models.Result)
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
	pagination := utility.ParsePagination(r)
	tracks, err := handler.albumsUsecase.GetFavoriteAlbums(userID, pagination)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, albumModels.MarshalAlbums)
}

//easyjson:json
type AlbumWithExtraInform struct {
	AlbumID     int                         `json:"album_id"`
	Tittle      string                      `json:"tittle"`
	Picture     string                      `json:"picture"`
	ReleaseDate string                      `json:"release_date"`
	Musician    *[]musiciansModels.Musician `json:"musician"`
	Tracks      []*models0.Track            `json:"tracks"`
}

func ConvertAlumToFullAlbum(album *albumModels.Album) *AlbumWithExtraInform {
	return &AlbumWithExtraInform{
		AlbumID:     album.AlbumID,
		Tittle:      album.Tittle,
		Picture:     album.Picture,
		ReleaseDate: album.ReleaseDate,
		Musician:    nil,
		Tracks:      nil,
	}
}

func MarshalAlbumWithExtraInform(data interface{}) ([]byte, error) {
	album, ok := data.(*AlbumWithExtraInform)
	if !ok {
		return nil, errors.New("cant convernt interface{} to album")
	}
	body, err := album.MarshalJSON()
	return body, err
}
