package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/tracks"
	tracksModels "2021_1_Noskool_team/internal/app/tracks/models"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"2021_1_Noskool_team/internal/pkg/utility"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type TracksHandler struct {
	router         *mux.Router
	tracksUsecase  tracks.Usecase
	logger         *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewTracksHandler(r *mux.Router, config *configs.Config, usecase tracks.Usecase) *TracksHandler {
	grpcCon, err := grpc.Dial(config.SessionMicroserviceAddr, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
	}

	handler := &TracksHandler{
		router:         r,
		tracksUsecase:  usecase,
		logger:         logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}
	err = ConfigLogger(handler, config)
	if err != nil {
		logrus.Error(err)
	}

	handler.router.Use(middleware.ContentTypeJson)
	authMiddleware := middleware.NewSessionMiddleware(handler.sessionsClient)

	handler.router.HandleFunc("/top",
		authMiddleware.CheckSessionMiddleware(handler.GetTop20Tracks)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/top/notauth",
		handler.GetTop20Tracks).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/toptrack",
		authMiddleware.CheckSessionMiddleware(handler.GetTopTrack)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/toptrack/notauth",
		handler.GetTopTrack).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/billbord",
		authMiddleware.CheckSessionMiddleware(handler.GetBillbordTopCharts)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/billbord/notauth",
		handler.GetBillbordTopChartsNotAuth).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}",
		handler.GetTrackByIDHandler)
	handler.router.HandleFunc("/mediateka",
		authMiddleware.CheckSessionMiddleware(handler.GetMediatekaForUser)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/favorites",
		authMiddleware.CheckSessionMiddleware(handler.GetFavoriteTracks)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/history",
		authMiddleware.CheckSessionMiddleware(handler.GetHistory)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}/history",
		authMiddleware.CheckSessionMiddleware(handler.AddToHistory)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/musician/{musician_id:[0-9]+}",
		authMiddleware.CheckSessionMiddleware(handler.GetTracksByMusicinID)).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}/picture",
		handler.UploadTrackPictureHandler).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}/audio",
		handler.UploadTrackAudioHandler).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}/favorite",
		authMiddleware.CheckSessionMiddleware(handler.AddDeleteTrackToFavorite)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/{track_id:[0-9]+}/mediateka",
		authMiddleware.CheckSessionMiddleware(handler.AddDeleteTrackToMediateka)).Methods(http.MethodPost, http.MethodOptions)
	handler.router.HandleFunc("/album/{album_id:[0-9]+}",
		handler.GetTracksByAlbumIDHandler).Methods(http.MethodGet, http.MethodOptions)
	handler.router.HandleFunc("/genre/{genre_id:[0-9]+}",
		handler.GetTracksByGenreIDHandler).Methods(http.MethodGet, http.MethodOptions)

	return handler
}

func (handler *TracksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func ConfigLogger(handler *TracksHandler, config *configs.Config) error {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}

	handler.logger.SetLevel(level)
	return nil
}

func (handler *TracksHandler) GetTrackByIDHandler(w http.ResponseWriter, r *http.Request) {
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct track id",
		})
		return
	}

	track, err := handler.tracksUsecase.GetTrackByID(trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}

	response.SendCorrectResponse(w, track, http.StatusOK, tracksModels.MarshalTrack)
}

func (handler *TracksHandler) UploadTrackPictureHandler(w http.ResponseWriter, r *http.Request) {
	trackID := mux.Vars(r)["track_id"]

	fileName, err := utility.SaveFile(r, "track_picture", "/static/img/tracks/", trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	fileNetPath := "/api/v1/data/img/tracks/" + *fileName
	trackIDINT, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	err = handler.tracksUsecase.UploadPicture(trackIDINT, fileNetPath)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *TracksHandler) UploadTrackAudioHandler(w http.ResponseWriter, r *http.Request) {
	trackID := mux.Vars(r)["track_id"]

	fileName, err := utility.SaveFile(r, "track_audio", "/static/audio/", trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	fileNetPath := "/api/v1/data/audio/" + *fileName
	trackIDINT, err := strconv.Atoi(trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}

	err = handler.tracksUsecase.UploadAudio(trackIDINT, fileNetPath)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusInternalServerError)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *TracksHandler) GetTracksByTittle(w http.ResponseWriter, r *http.Request) {
	trackTittle := mux.Vars(r)["track_tittle"]

	track, err := handler.tracksUsecase.GetTracksByTittle(trackTittle)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, track, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetTracksByMusicinID(w http.ResponseWriter, r *http.Request) {
	musicianID, err := strconv.Atoi(mux.Vars(r)["musician_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}

	track, err := handler.tracksUsecase.GetTrackByMusicianID(musicianID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, track, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetMediatekaForUser(w http.ResponseWriter, r *http.Request) {
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

	tracks, err := handler.tracksUsecase.GetTracksByUserID(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	handler.CheckTracksFavoriteAndMediateka(userID, tracks)

	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetFavoriteTracks(w http.ResponseWriter, r *http.Request) {
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
	tracks, err := handler.tracksUsecase.GetFavoriteTracks(userID, pagination)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	handler.CheckTracksFavoriteAndMediateka(userID, tracks)

	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) AddDeleteTrackToFavorite(w http.ResponseWriter, r *http.Request) {
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
	musicianID, err := strconv.Atoi(mux.Vars(r)["track_id"])
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
		err = handler.tracksUsecase.AddTrackToFavorites(userID, musicianID)
	} else if addOrDelete == "delete" {
		err = handler.tracksUsecase.DeleteTrackFromFavorites(userID, musicianID)
	}
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *TracksHandler) GetTracksByAlbumIDHandler(w http.ResponseWriter, r *http.Request) {
	albumID, err := strconv.Atoi(mux.Vars(r)["album_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}

	tracks, err := handler.tracksUsecase.GetTracksByAlbumID(albumID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetTracksByGenreIDHandler(w http.ResponseWriter, r *http.Request) {
	genreID, err := strconv.Atoi(mux.Vars(r)["genre_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}

	tracks, err := handler.tracksUsecase.GetTracksByGenreID(genreID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) AddDeleteTrackToMediateka(w http.ResponseWriter, r *http.Request) {
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
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	addOrDelete := r.URL.Query().Get("type")
	err = handler.tracksUsecase.AddDeleteTrackToMediateka(userID, trackID, addOrDelete)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *TracksHandler) GetTop20Tracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetTop20Tracks()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	session, ok := r.Context().Value("user_id").(models.Result)
	if ok {
		userID, err := strconv.Atoi(session.ID)
		if err == nil {
			handler.CheckTracksFavoriteAndMediateka(userID, tracks)
		}
	}

	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetTop20TracksNotAuth(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetTop20Tracks()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}

	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetTopTrack(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetTopTrack()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	session, ok := r.Context().Value("user_id").(models.Result)
	if ok {
		userID, err := strconv.Atoi(session.ID)
		if err == nil {
			handler.CheckTracksFavoriteAndMediateka(userID, tracks)
		}
	}

	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetTopTrackNotAuth(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetTopTrack()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetBillbordTopCharts(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetBillbordTopCharts()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	session, ok := r.Context().Value("user_id").(models.Result)
	if ok {
		userID, err := strconv.Atoi(session.ID)
		if err == nil {
			handler.CheckTracksFavoriteAndMediateka(userID, tracks)
		}
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetBillbordTopChartsNotAuth(w http.ResponseWriter, r *http.Request) {
	tracks, err := handler.tracksUsecase.GetBillbordTopCharts()
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	session, ok := r.Context().Value("user_id").(models.Result)
	if ok {
		userID, err := strconv.Atoi(session.ID)
		if err == nil {
			handler.CheckTracksFavoriteAndMediateka(userID, tracks)
		}
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
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
	tracks, err := handler.tracksUsecase.GetHistory(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK, tracksModels.MarshalTracks)
}

func (handler *TracksHandler) AddToHistory(w http.ResponseWriter, r *http.Request) {
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
	trackID, err := strconv.Atoi(mux.Vars(r)["track_id"])
	if err != nil {
		handler.logger.Error(err)
		response.SendErrorResponse(w, &commonModels.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Not correct musician id",
		})
		return
	}
	err = handler.tracksUsecase.AddToHistory(userID, trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendEmptyBody(w, http.StatusOK)
}

func (handler *TracksHandler) CheckTracksFavoriteAndMediateka(userID int, tracks []*tracksModels.Track) {
	for _, track := range tracks {
		handler.CheckTrackFlags(userID, track)
	}
}

func (handler *TracksHandler) CheckTrackFlags(userID int, track *tracksModels.Track) {
	track.InMediateka = handler.tracksUsecase.CheckTrackInMediateka(userID, track.TrackID)
	track.InFavorite = handler.tracksUsecase.CheckTrackInFavorite(userID, track.TrackID)
}
