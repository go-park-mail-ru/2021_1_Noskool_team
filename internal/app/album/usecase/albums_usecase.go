package usecase

import (
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/album/models"
)

type AlbumsUsecase struct {
	albumsRep album.Repository
}

func NewAlbumcUsecase(albumRep album.Repository) *AlbumsUsecase {
	return &AlbumsUsecase{
		albumsRep: albumRep,
	}
}

func (usecase *AlbumsUsecase) GetAlbumByID(albumID int) (*models.Album, error) {
	track, err := usecase.albumsRep.GetAlbumByID(albumID)
	return track, err
}
