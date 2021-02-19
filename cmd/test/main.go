package main

import (
	"2021_1_Noskool_team/internal/app/music/usecase"
	"2021_1_Noskool_team/internal/app/music/delivery/http"
	"2021_1_Noskool_team/internal/pkg/server"
	grpcSerc "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/server"
	sesUsecase "2021_1_Noskool_team/internal/microservices/auth/usecase"
	"fmt"
)

func main() {
	sessionsUsecase := sesUsecase.NewSessionsUsecase()
	go grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, "127.0.0.1:8081")


	config := server.Config{BindAddr: ":8080"}
	musicUsecase := usecase.MusicUsecase{}
	handler := http.NewMusicHandler(musicUsecase)
	err := server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
}
