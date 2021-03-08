package main

import (
	"2021_1_Noskool_team/configs"
	grpcSerc "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/server"
	sesUsecase "2021_1_Noskool_team/internal/microservices/auth/usecase"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	configPath = "configs/config.toml"
)

func main() {
	time.Sleep(5 * time.Second)
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	sessionsUsecase := sesUsecase.NewSessionsUsecase(config)
	grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, config.SessionMicroserviceAddr)
}
