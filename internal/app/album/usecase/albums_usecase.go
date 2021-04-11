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
	album, err := usecase.albumsRep.GetAlbumByID(albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (usecase *AlbumsUsecase) GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error) {
	album, err := usecase.albumsRep.GetAlbumsByMusicianID(musicianID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (usecase *AlbumsUsecase) GetAlbumsByTrackID(trackID int) (*[]models.Album, error) {
	album, err := usecase.albumsRep.GetAlbumsByMusicianID(trackID)
	if err != nil {
		return nil, err
	}
	return album, nil
}
