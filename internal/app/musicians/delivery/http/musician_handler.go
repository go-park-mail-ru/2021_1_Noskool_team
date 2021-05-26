package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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

	authmiddlware := middleware.NewSessionMiddleware(handler.sessionsClient)

	handler.router.HandleFunc("/top", authmiddlware.CheckSessionMiddleware(handler.GetMusicians)).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/top/notauth", handler.GetMusicians).Methods("GET", http.MethodOptions)

	handler.router.HandleFunc("/genres", handler.GetGenreForMusician).Methods("GET", http.MethodOptions)

	handler.router.HandleFunc("/popular", authmiddlware.CheckSessionMiddleware(handler.GetMusiciansTop4)).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bygenre/{genre}", handler.GetMusiciansByGenre).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/{musician_id:[0-9]+}",
		authmiddlware.CheckSessionMiddleware(handler.GetMusicianByID)).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/bytrack/{track_id:[0-9]+}", handler.GetMusicianByTrackID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/byalbum/{album_id:[0-9]+}", handler.GetMusicianByAlbumID).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/byplaylist/{playlist_id:[0-9]+}", handler.GetMusicianByPlaylistID).Methods("GET", http.MethodOptions)

	handler.router.HandleFunc("/mediateka", authmiddlware.CheckSessionMiddleware(handler.GetMediatekaForUser)).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/favorites", authmiddlware.CheckSessionMiddleware(handler.GetFavoritesForUser)).Methods("GET", http.MethodOptions)
	handler.router.HandleFunc("/{musician_id:[0-9]+}/mediateka",
		authmiddlware.CheckSessionMiddleware(handler.AddDeleteMusicianToMediateka)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{musician_id:[0-9]+}/favorites",
		authmiddlware.CheckSessionMiddleware(handler.AddDeleteMusicianToFavorites)).Methods(http.MethodPost, http.MethodOptions)

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

func (handler *MusiciansHandler) GetGenreForMusician(w http.ResponseWriter, r *http.Request) {
	type MusicianName struct {
		Name string `json:"name"`
	}
	req := &MusicianName{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}

	geners, err := handler.musicianUsecase.GetGenreForMusician(req.Name)
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	fmt.Println(geners)
	response.SendCorrectResponse(w, geners, 200, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusiciansByGenre(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	genre, ok := vars["genre"]
	if !ok {
		handler.logger.Errorf("Error get genre from query string")
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusiciansByGenre(genre)
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansByGenres: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusicianByID(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	musicianID, ok := vars["musician_id"]
	if !ok {
		handler.logger.Errorf("Error get musician_id from query string")
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	musicianIDint, err := strconv.Atoi(musicianID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	musician, err := handler.musicianUsecase.GetMusicianByID(musicianIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	fullMusician := &musiciansModels.MusicianFullInformation{
		MusicianID:  musician.MusicianID,
		Name:        musician.Name,
		Description: musician.Description,
		Picture:     musician.Picture,
		InMediateka: false,
		InFavorite:  false,
	}
	err = handler.musicianUsecase.CheckMusicianInMediateka(userID, musician.MusicianID)
	if err == nil {
		fullMusician.InMediateka = true
	}
	err = handler.musicianUsecase.CheckMusicianInFavorite(userID, musician.MusicianID)
	if err == nil {
		fullMusician.InFavorite = true
	}

	response.SendCorrectResponse(w, fullMusician, 200, musiciansModels.MarshalMusicianFullInform)
}

func (handler *MusiciansHandler) GetMusicianByTrackID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	trackID, ok := vars["track_id"]
	if !ok {
		handler.logger.Errorf("Error get track_id from query string")
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	trackIDint, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusicianByTrackID(trackIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByTrackID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusicianByAlbumID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
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
	musicians, err := handler.musicianUsecase.GetMusicianByAlbumID(albumIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByAlbumID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusicianByPlaylistID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	playlistID, ok := vars["playlist_id"]
	if !ok {
		handler.logger.Errorf("Error get playlist_id from query string")
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	playlistIDint, err := strconv.Atoi(playlistID)
	if err != nil {
		handler.logger.Error(err)
		_, _ = w.Write(response.FailedResponse(w, 400))
		return
	}
	musicians, err := handler.musicianUsecase.GetMusicianByPlaylistID(playlistIDint)
	if err != nil {
		handler.logger.Errorf("Error in GetMusicianByPlaylistID: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, 200, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusiciansTop4(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	musicians, err := handler.musicianUsecase.GetMusiciansTop4()
	if err != nil {
		handler.logger.Errorf("Error in GetMusiciansTop3: %v", err)
		_, _ = w.Write(response.FailedResponse(w, 500))
		return
	}
	response.SendCorrectResponse(w, musicians, http.StatusOK, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetMusicians(w http.ResponseWriter, r *http.Request) {
	musicians, err := handler.musicianUsecase.GetMusicians()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, musicians, http.StatusOK, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) AddDeleteMusicianToMediateka(w http.ResponseWriter, r *http.Request) {
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
	musicianID, err := strconv.Atoi(mux.Vars(r)["musician_id"])
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
		err = handler.musicianUsecase.AddMusicianToMediateka(userID, musicianID)
	} else {
		err = handler.musicianUsecase.DeleteMusicianFromMediateka(userID, musicianID)
	}
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *MusiciansHandler) AddDeleteMusicianToFavorites(w http.ResponseWriter, r *http.Request) {
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
	musicianID, err := strconv.Atoi(mux.Vars(r)["musician_id"])
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
		err = handler.musicianUsecase.AddMusicianToFavorites(userID, musicianID)
	} else {
		err = handler.musicianUsecase.DeleteMusicianFromFavorites(userID, musicianID)
	}
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *MusiciansHandler) GetMediatekaForUser(w http.ResponseWriter, r *http.Request) {
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

	musicians, err := handler.musicianUsecase.GetMusiciansMediateka(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	fullMusicians := make([]*musiciansModels.MusicianFullInformation, 0)
	for _, item := range musicians {
		newMusician := &musiciansModels.MusicianFullInformation{
			MusicianID:  item.MusicianID,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			InMediateka: false,
			InFavorite:  false,
		}

		err = handler.musicianUsecase.CheckMusicianInMediateka(userID, item.MusicianID)
		if err == nil {
			newMusician.InMediateka = true
		}
		err = handler.musicianUsecase.CheckMusicianInFavorite(userID, item.MusicianID)
		if err == nil {
			newMusician.InFavorite = true
		}
		fullMusicians = append(fullMusicians, newMusician)
	}

	response.SendCorrectResponse(w, fullMusicians, http.StatusOK, musiciansModels.MarshalMusicians)
}

func (handler *MusiciansHandler) GetFavoritesForUser(w http.ResponseWriter, r *http.Request) {
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

	musicians, err := handler.musicianUsecase.GetMusiciansFavorites(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}

	fullMusicians := make([]*musiciansModels.MusicianFullInformation, 0)
	for _, item := range musicians {
		newMusician := &musiciansModels.MusicianFullInformation{
			MusicianID:  item.MusicianID,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			InMediateka: false,
			InFavorite:  false,
		}

		err = handler.musicianUsecase.CheckMusicianInMediateka(userID, item.MusicianID)
		if err == nil {
			newMusician.InMediateka = true
		}
		err = handler.musicianUsecase.CheckMusicianInFavorite(userID, item.MusicianID)
		if err == nil {
			newMusician.InFavorite = true
		}
		fullMusicians = append(fullMusicians, newMusician)
	}

	response.SendCorrectResponse(w, fullMusicians, http.StatusOK, musiciansModels.MarshalMusicians)
}
