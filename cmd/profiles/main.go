package main

import (
	"2021_1_Noskool_team/configs"
	profiles "2021_1_Noskool_team/internal/app/profiles/delivery/http"
	grpcSerc "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/server"
	sesUsecase "2021_1_Noskool_team/internal/microservices/auth/usecase"
	"flag"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := profiles.New(config)
	go s.Start()

	// if err := s.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	sessionsUsecase := sesUsecase.NewSessionsUsecase(config)
	grpcSerc.StartSessionsGRPCServer(&sessionsUsecase, config.SessionMicroserviceAddr)
}
