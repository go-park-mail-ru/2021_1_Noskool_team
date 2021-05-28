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
	"errors"
	"fmt"
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
		authmiddlware.CheckSessionMiddleware(handler.GetFavoriteAlbums)).Methods(http.MethodGet, http.MethodOptions)

	handler.router.HandleFunc("/mediateka",
		authmiddlware.CheckSessionMiddleware(handler.GetAlbumsMediateka)).Methods(http.MethodGet, http.MethodOptions)

	handler.router.HandleFunc("/top",
		authmiddlware.CheckSessionMiddleware(handler.GetAlbums)).Methods("GET", http.MethodOptions)

	handler.router.HandleFunc("/top/notauth", handler.GetAlbums).Methods("GET", http.MethodOptions)

	handler.router.HandleFunc("/{album_id:[0-9]+}", handler.GetAlbumByID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bymusician/{musician_id:[0-9]+}", handler.GetAlbumsByMusicianID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bytrack/{track_id:[0-9]+}", handler.GetAlbumsByTrackID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/{album_id:[0-9]+}/favorites",
		authmiddlware.CheckSessionMiddleware(handler.AddDeleteAlbumToFavorites)).Methods("POST", http.MethodOptions)
	handler.router.HandleFunc("/{album_id:[0-9]+}/mediateka",
		authmiddlware.CheckSessionMiddleware(handler.AddDeleteAlbumToMediateka)).Methods("POST", http.MethodOptions)

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
	vars := mux.Vars(r)
	albumID, ok := vars["album_id"]
	if !ok {
		handler.logger.Errorf("Error get album_id from query string")
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	albumIDint, err := strconv.Atoi(albumID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	albumByID, err := handler.albumsUsecase.GetAlbumByID(albumIDint)
	albumWithTracks := ConvertAlumToFullAlbum(albumByID)
	albumWithTracks.Tracks, _ = handler.tracksUsecase.GetTracksByAlbumID(albumWithTracks.AlbumID)
	albumWithTracks.Musician, _ = handler.musUsecase.GetMusicianByAlbumID(albumWithTracks.AlbumID)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
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
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	musicianIDint, err := strconv.Atoi(musicianID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	albums, err := handler.albumsUsecase.GetAlbumsByMusicianID(musicianIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumsByMusicianID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, albums, 200, albumModels.MarshalAlbums)
}

func (handler *AlbumsHandler) GetAlbumsByTrackID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	trackID, ok := vars["track_id"]
	if !ok {
		handler.logger.Errorf("Error get album_id from query string")
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	trackIDint, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	albums, err := handler.albumsUsecase.GetAlbumsByTrackID(trackIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetAlbumByTrackID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, albums, 200, albumModels.MarshalAlbums)
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

func (handler *AlbumsHandler) GetAlbumsMediateka(w http.ResponseWriter, r *http.Request) {
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
	favoriteAlbums, err := handler.albumsUsecase.GetAlbumsMediateka(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	albumsFullInf := make([]*albumModels.AlbumFullInformation, 0)
	for _, album := range favoriteAlbums {
		newAlbum := &albumModels.AlbumFullInformation{
			AlbumID:     album.AlbumID,
			Tittle:      album.Tittle,
			Picture:     album.Picture,
			ReleaseDate: album.ReleaseDate,
		}
		err = handler.albumsUsecase.CheckAlbumInMediateka(userID, album.AlbumID)
		if err == nil {
			newAlbum.InMediateka = true
		}
		err = handler.albumsUsecase.CheckAlbumInFavorite(userID, album.AlbumID)
		if err == nil {
			newAlbum.InFavorite = true
		}
		albumsFullInf = append(albumsFullInf, newAlbum)
	}

	response.SendCorrectResponse(w, albumsFullInf, http.StatusOK, albumModels.MarshalAlbums)
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
	favoriteAlbums, err := handler.albumsUsecase.GetFavoriteAlbums(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	albumsFullInf := make([]*albumModels.AlbumFullInformation, 0)
	for _, album := range favoriteAlbums {
		newAlbum := &albumModels.AlbumFullInformation{
			AlbumID:     album.AlbumID,
			Tittle:      album.Tittle,
			Picture:     album.Picture,
			ReleaseDate: album.ReleaseDate,
		}
		err = handler.albumsUsecase.CheckAlbumInMediateka(userID, album.AlbumID)
		if err == nil {
			newAlbum.InMediateka = true
		}
		err = handler.albumsUsecase.CheckAlbumInFavorite(userID, album.AlbumID)
		if err == nil {
			newAlbum.InFavorite = true
		}
		albumsFullInf = append(albumsFullInf, newAlbum)
	}

	response.SendCorrectResponse(w, albumsFullInf, http.StatusOK, albumModels.MarshalAlbums)
}

func (handler *AlbumsHandler) GetAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	albums, err := handler.albumsUsecase.GetAlbums()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, albums, http.StatusOK, albumModels.MarshalAlbums)
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
	albumExtra, ok := data.(*AlbumWithExtraInform)
	if !ok {
		return nil, errors.New("cant convernt interface{} to album")
	}
	body, err := albumExtra.MarshalJSON()
	return body, err
}
