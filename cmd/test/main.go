package main

import (
	"2021_1_Noskool_team/configs"
	musicHttp "2021_1_Noskool_team/internal/app/music/delivery/http"
	musicianUsecase "2021_1_Noskool_team/internal/app/musicans/usecase"
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

	musUsecase := musicianUsecase.NewMusicsUsecase(config)
	trackUse := trackUsecase.NewTracksUsecase(config)
	handler := musicHttp.NewFinalHandler(config, &trackUse, &musUsecase)
	err = server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
}
