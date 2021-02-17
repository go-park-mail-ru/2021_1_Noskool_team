package http

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"testWorkWithAuth/internal/app/music"
)

type MusicHandler struct {
	router       *mux.Router
	musicUsecase *music.Usecase
	logger       *logrus.Logger
}

func NewMusicHandler(usecase music.Usecase) *MusicHandler {
	handler := &MusicHandler{
		router:       mux.NewRouter(),
		musicUsecase: &usecase,
		logger:       logrus.New(),
	}

	handler.router.HandleFunc("/getMusic", handler.GetMusic)
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	return handler
}

func (handler *MusicHandler) GetMusic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Music"))
}

func (handler *MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
