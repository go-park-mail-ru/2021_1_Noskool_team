package http

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	albumHttp "2021_1_Noskool_team/internal/app/album/delivery/http"
	"2021_1_Noskool_team/internal/app/middleware"
	"2021_1_Noskool_team/internal/app/musicians"
	musicHttp "2021_1_Noskool_team/internal/app/musicians/delivery/http"
	"2021_1_Noskool_team/internal/app/playlists"
	playlistHttp "2021_1_Noskool_team/internal/app/playlists/delivery/http"
	"2021_1_Noskool_team/internal/app/search"
	searchHttp "2021_1_Noskool_team/internal/app/search/delivery/http"
	"2021_1_Noskool_team/internal/app/tracks"
	trackHttp "2021_1_Noskool_team/internal/app/tracks/delivery/http"
	"2021_1_Noskool_team/internal/pkg/monitoring"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
)

type MusicHandler struct {
	router          *mux.Router
	tracksHandler   *trackHttp.TracksHandler
	musicianHandler *musicHttp.MusiciansHandler
	albumsHandler   *albumHttp.AlbumsHandler
	playlistHandler *playlistHttp.PlaylistsHandler
	searchHandler   *searchHttp.SearchHandler
}

func (handler MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}

func NewFinalHandler(config *configs.Config, tracksUsecase tracks.Usecase,
	musicUsecase musicians.Usecase, albumsUsecase album.Usecase,
	playlistUsecase playlists.Usecase, searchUsecase search.Usecase) *MusicHandler {
	handler := &MusicHandler{
		router: mux.NewRouter(),
	}

	metricks := monitoring.RegisterMetrics(handler.router)

	logrus.Info(config.MediaFolder)

	handler.router.PathPrefix("/api/v1/music/data/").
		Handler(
			http.StripPrefix(
				"/api/v1/music/data/", http.FileServer(http.Dir(config.MediaFolder))))

	sanitizer := bluemonday.UGCPolicy()

	handler.router.Use(middleware.LoggingMiddleware(metricks))

	musicRouter := handler.router.PathPrefix("/api/v1/music/musician/").Subrouter()
	tracksRouter := handler.router.PathPrefix("/api/v1/music/track/").Subrouter()
	albumsRouter := handler.router.PathPrefix("/api/v1/music/album/").Subrouter()
	searchRouter := handler.router.PathPrefix("/api/v1/music/search/").Subrouter()
	playlistsRouter := handler.router.PathPrefix("/api/v1/music/playlist/").Subrouter()
	handler.musicianHandler = musicHttp.NewMusicHandler(musicRouter, config, musicUsecase)
	handler.tracksHandler = trackHttp.NewTracksHandler(tracksRouter, config, tracksUsecase)
	handler.albumsHandler = albumHttp.NewAlbumsHandler(albumsRouter, config, albumsUsecase,
		tracksUsecase, musicUsecase)
	handler.searchHandler = searchHttp.NewSearchHandler(searchRouter, config, searchUsecase, sanitizer)
	handler.playlistHandler = playlistHttp.NewPlaylistsHandler(playlistsRouter, config, playlistUsecase,
		albumsUsecase)

	handler.router.HandleFunc("/api/v1/music/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("main main page"))
	})

	CORSMiddleware := middleware.NewCORSMiddleware(config)
	handler.router.Use(CORSMiddleware.CORS)
	handler.router.Use(middleware.PanicMiddleware(metricks))
	return handler
}
