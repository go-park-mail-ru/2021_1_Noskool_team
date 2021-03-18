package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	commonModels "2021_1_Noskool_team/internal/models"
	"2021_1_Noskool_team/internal/pkg/response"
	"2021_1_Noskool_team/internal/pkg/utility"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
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

	//handler.router.Use(middleware.ContentTypeJson)

	handler.router.HandleFunc("/{track_id:[0-9]+}",
		middleware.ContentTypeJson(handler.GetTrackByIDHandler))
	handler.router.HandleFunc("/",
		middleware.ContentTypeJson(handler.GetTracksByUserID)).Methods(http.MethodGet)
	handler.router.HandleFunc("/favorites",
		middleware.ContentTypeJson(handler.GetFavoriteTracks)).Methods(http.MethodGet)
	handler.router.HandleFunc("/{track_tittle}",
		middleware.ContentTypeJson(handler.GetTracksByTittle)).Methods(http.MethodGet)
	handler.router.HandleFunc("/musician/{musician_id:[0-9]+}",
		middleware.ContentTypeJson(handler.GetTrackByMusicianID)).Methods(http.MethodGet)
	handler.router.HandleFunc("/{track_id:[0-9]+}/picture",
		handler.UploadTrackPictureHandler).Methods(http.MethodPost)
	handler.router.HandleFunc("/{track_id:[0-9]+}/audio",
		handler.UploadTrackAudioHandler).Methods(http.MethodPost)

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
	response.SendCorrectResponse(w, track, http.StatusOK)
}

func (handler *TracksHandler) UploadTrackPictureHandler(w http.ResponseWriter, r *http.Request) {
	trackID := mux.Vars(r)["track_id"]

	fileName, err := utility.SaveFile(r, "track_picture", "/static/img/tracks/", trackID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusBadRequest)
		return
	}
	fileNetPath := "/api/v1/data/img/track/" + *fileName
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
	fileNetPath := "/api/v1/data/audio/track/" + *fileName
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
	response.SendCorrectResponse(w, track, http.StatusOK)
}

func (handler *TracksHandler) GetTrackByMusicianID(w http.ResponseWriter, r *http.Request) {
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
	response.SendCorrectResponse(w, track, http.StatusOK)
}

func (handler *TracksHandler) GetTracksByUserID(w http.ResponseWriter, r *http.Request) {
	SessionHash, _ := r.Cookie("session_id")
	session, err := handler.sessionsClient.Check(context.Background(), SessionHash.Value)
	if err != nil {
		handler.logger.Error(err)
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
	response.SendCorrectResponse(w, tracks, http.StatusOK)
}

func (handler *TracksHandler) GetFavoriteTracks(w http.ResponseWriter, r *http.Request) {
	SessionHash, _ := r.Cookie("session_id")
	session, err := handler.sessionsClient.Check(context.Background(), SessionHash.Value)
	if err != nil {
		handler.logger.Error(err)
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

	tracks, err := handler.tracksUsecase.GetFavoriteTracks(userID)
	if err != nil {
		handler.logger.Error(err)
		response.SendEmptyBody(w, http.StatusNoContent)
		return
	}
	response.SendCorrectResponse(w, tracks, http.StatusOK)
}
