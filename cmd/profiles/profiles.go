package main

import (
	"2021_1_Noskool_team/configs"
	profiles "2021_1_Noskool_team/internal/app/profiles/delivery/http"
	"2021_1_Noskool_team/internal/app/profiles/repository/postgresDB"
	"2021_1_Noskool_team/internal/app/profiles/usecase"
	"2021_1_Noskool_team/internal/pkg/utility"
	"flag"
	"log"
	"time"

	"github.com/BurntSushi/toml"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	time.Sleep(50 * time.Second)
	flag.Parse()

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	profDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	profRep := postgresDB.NewProfileRepository(profDBCon)
	profUsecase := usecase.NewProfilesUsecase(profRep)
	sanitizer := bluemonday.UGCPolicy()

	s := profiles.New(config, profUsecase, sanitizer)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
