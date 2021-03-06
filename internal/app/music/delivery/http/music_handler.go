package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	albumHttp "2021_1_Noskool_team/internal/app/album/delivery/http"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	musicHttp "2021_1_Noskool_team/internal/app/musicians/delivery/http"
	"2021_1_Noskool_team/internal/app/tracks"
	trackHttp "2021_1_Noskool_team/internal/app/tracks/delivery/http"
	"github.com/gorilla/mux"
	"net/http"
)

type MusicHandler struct {
	router          *mux.Router
	tracksHandler   *trackHttp.TracksHandler
	musicianHandler *musicHttp.MusiciansHandler
	albumsHandler   *albumHttp.AlbumsHandler
}

func (handler MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func NewFinalHandler(config *configs.Config, tracksUsecase tracks.Usecase,
	musicUsecase musicians.Usecase, albumsUsecase album.Usecase) *MusicHandler {
	handler := &MusicHandler{
		router: mux.NewRouter(),
	}
	musicRouter := handler.router.PathPrefix("/api/v1/musician/").Subrouter()
	tracksRouter := handler.router.PathPrefix("/api/v1/track/").Subrouter()
	albumsRouter := handler.router.PathPrefix("/api/v1/album/").Subrouter()
	handler.musicianHandler = musicHttp.NewMusicHandler(musicRouter, config, musicUsecase)
	handler.tracksHandler = trackHttp.NewTracksHandler(tracksRouter, config, tracksUsecase)
	handler.albumsHandler = albumHttp.NewAlbumsHandler(albumsRouter, config, albumsUsecase)

	handler.router.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main main page"))
	})
	handler.router.Use(middleware.LoggingMiddleware)

	CORSMiddleware := middleware.NewCORSMiddleware(config)
	handler.router.Use(CORSMiddleware.CORS)
	handler.router.Use(middleware.PanicMiddleware)
	return handler
}
