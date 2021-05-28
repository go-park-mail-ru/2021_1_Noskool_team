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
	album, err := usecase.albumsRep.GetAlbumsByTrackID(trackID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (usecase *AlbumsUsecase) SearchAlbums(searchQuery string) ([]*models.Album, error) {
	album, err := usecase.albumsRep.SearchAlbums(searchQuery)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (usecase *AlbumsUsecase) AddAlbumToFavorites(userID, albumID int) error {
	err := usecase.albumsRep.CheckAlbumInMediateka(userID, albumID)
	if err != nil {
		err = usecase.albumsRep.AddAlbumToMediateka(userID, albumID)
		if err != nil {
			return err
		}
	}
	err = usecase.albumsRep.AddAlbumToFavorites(userID, albumID)
	return err
}

func (usecase *AlbumsUsecase) DelteAlbumFromFavorites(userID, albumID int) error {
	err := usecase.albumsRep.DelteAlbumFromFavorites(userID, albumID)
	return err
}

func (usecase *AlbumsUsecase) AddAlbumToMediateka(userID, albumID int) error {
	err := usecase.albumsRep.AddAlbumToMediateka(userID, albumID)
	return err
}

func (usecase *AlbumsUsecase) DeleteAlbumFromMediateka(userID, albumID int) error {
	err := usecase.albumsRep.DeleteAlbumFromMediateka(userID, albumID)
	return err
}

func (usecase *AlbumsUsecase) GetFavoriteAlbums(userID int) ([]*models.Album, error) {
	albums, err := usecase.albumsRep.GetFavoriteAlbums(userID)
	return albums, err
}

func (usecase *AlbumsUsecase) GetAlbums() ([]*models.Album, error) {
	albums, err := usecase.albumsRep.GetAlbums()
	return albums, err
}

func (usecase *AlbumsUsecase) CheckAlbumInFavorite(userID, albumID int) error {
	return usecase.albumsRep.CheckAlbumInFavorite(userID, albumID)
}

func (usecase *AlbumsUsecase) CheckAlbumInMediateka(userID, albumID int) error {
	return usecase.albumsRep.CheckAlbumInMediateka(userID, albumID)
}

func (usecase *AlbumsUsecase) GetAlbumsMediateka(userID int) ([]*models.Album, error) {
	albums, err := usecase.albumsRep.GetAlbumsMediateka(userID)
	return albums, err
}
