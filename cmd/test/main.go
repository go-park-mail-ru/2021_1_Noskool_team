package main

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/music/delivery/http"
	"2021_1_Noskool_team/internal/app/music/usecase"
	grpcSerc "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/server"
	sesUsecase "2021_1_Noskool_team/internal/microservices/auth/usecase"
	"2021_1_Noskool_team/internal/pkg/server"
	"fmt"
	"github.com/BurntSushi/toml"
)

const (
	configPath = "configs/config.toml"
)

func main() {
	fmt.Println("hello")
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		fmt.Println(err)
	}

	sessionsUsecase := sesUsecase.NewSessionsUsecase(config)
	go grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, config.SessionMicroserviceAddr)

	musicUsecase := usecase.MusicUsecase{}
	handler := http.NewMusicHandler(config, musicUsecase)
	err = server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
}
