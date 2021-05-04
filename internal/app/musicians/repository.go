package musicians

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
)

// mockgen -destination=mocks/repository_mock.go -source=repository.go
type Repository interface {
	SearchMusicians(searchQuery string) ([]*models.Musician, error)
	GetMusiciansByGenre(genre string) (*[]models.Musician, error)
	GetMusicianByID(musicianID int) (*models.Musician, error)

	GetMusicianByTrackID(trackID int) (*[]models.Musician, error)
	GetMusicianByAlbumID(albumID int) (*[]models.Musician, error)
	GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error)
	GetMusiciansTop4() (*[]models.Musician, error)
	GetMusicians() (*[]models.Musician, error)
}
