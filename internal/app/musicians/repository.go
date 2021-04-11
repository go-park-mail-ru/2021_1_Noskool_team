package musicians

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
)

type Repository interface {
	SearchMusicians(searchQuery string) ([]*models.Musician, error)
	GetMusiciansByGenres(genre string) (*[]models.Musician, error)
	GetMusicianByID(musicianID int) (*models.Musician, error)

	GetMusicianByTrackID(trackID int) (*[]models.Musician, error)
	GetMusicianByAlbumID(albumID int) (*[]models.Musician, error)
	GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error)
}
