package album

import "2021_1_Noskool_team/internal/app/album/models"

type Repository interface {
	GetAlbumByID(albumID int) (*models.Album, error)
	GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error)
	GetAlbumsByTrackID(trackID int) (*[]models.Album, error)
	SearchAlbums(searchQuery string) ([]*models.Album, error)
}
