package main

import (
	"2021_1_Noskool_team/configs"
	finalHttp "2021_1_Noskool_team/internal/app/final/delivery/http"
	musicUsecase "2021_1_Noskool_team/internal/app/music/usecase"
	trackUsecase "2021_1_Noskool_team/internal/app/tracks/usecase"
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
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		fmt.Println(err)
	}

	sessionsUsecase := sesUsecase.NewSessionsUsecase(config)
	go grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, config.SessionMicroserviceAddr)

	musUsecase := musicUsecase.NewMusicsUsecase(config)
	trackUse := trackUsecase.NewTracksUsecase(config)
	//handler := http.NewMusicHandler(config, &musUsecase)
	handler := finalHttp.NewFinalHandler(config, &trackUse, &musUsecase)
	err = server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
}
