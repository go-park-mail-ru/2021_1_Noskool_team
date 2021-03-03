package main

import (
	"2021_1_Noskool_team/configs"
	albumUsecase "2021_1_Noskool_team/internal/app/album/usecase"
	musicHttp "2021_1_Noskool_team/internal/app/music/delivery/http"
	musicianUsecase "2021_1_Noskool_team/internal/app/musicians/usecase"
	trackRepository "2021_1_Noskool_team/internal/app/tracks/repository"
	trackUsecase "2021_1_Noskool_team/internal/app/tracks/usecase"
	"2021_1_Noskool_team/internal/pkg/server"
	"2021_1_Noskool_team/internal/pkg/utility"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

const (
	configPath = "configs/config.toml"
)

func main() {
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	musUsecase := musicianUsecase.NewMusicsUsecase(config)

	tracDBCon, err := utility.CreatePostgresConnection(config.MusicPostgresBD)
	if err != nil {
		logrus.Error(err)
	}
	trackRep := trackRepository.NewTracksRepository(tracDBCon)

	trackUse := trackUsecase.NewTracksUsecase(trackRep)
	albumsUse := albumUsecase.NewAlbumcUsecase(config)
	handler := musicHttp.NewFinalHandler(config, trackUse, musUsecase, albumsUse)
	fmt.Println("Нормально запустились")
	err = server.Start(config, handler)
	if err != nil {
		fmt.Println("fail start server")
	}
	fmt.Println("Закончили работу")
}
