package album

import (
	"2021_1_Noskool_team/internal/app/album/models"
)

type Repository interface {
	GetAlbumByID(albumID int) (*models.Album, error)
	GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error)
	GetAlbumsByTrackID(trackID int) (*[]models.Album, error)
	SearchAlbums(searchQuery string) ([]*models.Album, error)
	AddAlbumToFavorites(userID, albumID int) error
	DelteAlbumFromFavorites(userID, albumID int) error
	AddAlbumToMediateka(userID, albumID int) error
	DeleteAlbumFromMediateka(userID, albumID int) error
	GetFavoriteAlbums(userID int) ([]*models.Album, error)
	CheckAlbumInMediateka(userID, albumID int) error
	GetAlbums() ([]*models.Album, error)
	CheckAlbumInFavorite(userID, albumID int) error
	GetAlbumsMediateka(userID int) ([]*models.Album, error)
}
