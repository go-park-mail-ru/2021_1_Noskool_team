package album

import (
	"2021_1_Noskool_team/internal/app/album/models"
	commonModels "2021_1_Noskool_team/internal/models"
)

type Usecase interface {
	GetAlbumByID(albumID int) (*models.Album, error)
	GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error)
	GetAlbumsByTrackID(trackID int) (*[]models.Album, error)
	SearchAlbums(searchQuery string) ([]*models.Album, error)
	AddAlbumToFavorites(userID, albumID int) error
	DelteAlbumFromFavorites(userID, albumID int) error
	AddAlbumToMediateka(userID, albumID int) error
	DeleteAlbumFromMediateka(userID, albumID int) error
	GetFavoriteAlbums(userID int, pagination *commonModels.Pagination) ([]*models.Album, error)
	GetAlbums() ([]*models.Album, error)
}
