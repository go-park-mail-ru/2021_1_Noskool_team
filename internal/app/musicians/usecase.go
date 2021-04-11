package musicians

import "2021_1_Noskool_team/internal/app/musicians/models"

// mockgen -destination=mocks/usecase_mock.go -source=usecase.go
type Usecase interface {
	GetMusiciansByGenre(genre string) (*[]models.Musician, error)
	GetMusicianByID(musicianID int) (*models.Musician, error)

	GetMusicianByTrackID(trackID int) (*[]models.Musician, error)
	GetMusicianByAlbumID(albumID int) (*[]models.Musician, error)
	GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error)
}
