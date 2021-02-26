package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/music"
	musicHttp "2021_1_Noskool_team/internal/app/music/delivery/http"
	"2021_1_Noskool_team/internal/app/tracks"
	trackHttp "2021_1_Noskool_team/internal/app/tracks/delivery/http"
	"github.com/gorilla/mux"
	"net/http"
)

type FinalHandler struct {
	router        *mux.Router
	tracksHandler *trackHttp.TracksHandler
	musicHandler  *musicHttp.MusicHandler
}

func (handler FinalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func NewFinalHandler(config *configs.Config, tracksUsecase tracks.Usecase, musicUsecase music.Usecase) *FinalHandler {

	handler := &FinalHandler{
		router:        mux.NewRouter(),

	}
	musicRouter := handler.router.PathPrefix("/music/").Subrouter()
	tracksRouter := handler.router.PathPrefix("/tracks/").Subrouter()
	handler.musicHandler = musicHttp.NewMusicHandler(musicRouter, config, musicUsecase)
	handler.tracksHandler = trackHttp.NewTracksHandler(tracksRouter, config, tracksUsecase)
	handler.router.Handle("/music/", handler.musicHandler)
	handler.router.Handle("/tracks/", handler.tracksHandler )

	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main main page"))
	})

	return handler
}
