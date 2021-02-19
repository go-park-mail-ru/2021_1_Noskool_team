package http

import (
	"2021_1_Noskool_team/internal/app/music"
	"2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type MusicHandler struct {
	router       *mux.Router
	musicUsecase *music.Usecase
	logger       *logrus.Logger
	sessionsClient client.AuthCheckerClient
}

func NewMusicHandler(usecase music.Usecase) *MusicHandler {
	grpcCon, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure(),)
	fmt.Println(err)

	handler := &MusicHandler{
		router:       mux.NewRouter(),
		musicUsecase: &usecase,
		logger:       logrus.New(),
		sessionsClient: client.NewSessionsClient(grpcCon),
	}

	handler.router.HandleFunc("/getMusic", handler.GetMusic)
	handler.router.HandleFunc("/createSession", handler.CreateSession)
	handler.router.HandleFunc("/getMusic", handler.GetMusic)
	handler.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	return handler
}

func (handler *MusicHandler) GetMusic(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Music"))
}

func (handler *MusicHandler) CreateSession(w http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(r.FormValue("user_id"))

	session, err := handler.sessionsClient.Create(context.Background(), userID)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   strconv.Itoa(session.ID),
		Expires: time.Now().Add(5 * time.Hour),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte(strconv.Itoa(session.ID)))
}

func (handler *MusicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.router.ServeHTTP(w, r)
}
