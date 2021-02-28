package usecase

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/album/models"
	"2021_1_Noskool_team/internal/app/album/repository"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type AlbumsUsecase struct {
	albumsRep album.Repository
}

func NewAlbumcUsecase(config *configs.Config) *AlbumsUsecase {
	dbCon, err := sql.Open("postgres",
		config.MusicPostgresBD,
	)
	if err != nil {
		logrus.Error(err)
	}

	return &AlbumsUsecase{
		albumsRep: repository.NewAlbumsRepository(dbCon),
	}
}

func (usecase *AlbumsUsecase) GetAlbumByID(albumID int) (*models.Album, error) {
	track, err := usecase.albumsRep.GetAlbumByID(albumID)
	return track, err
}
