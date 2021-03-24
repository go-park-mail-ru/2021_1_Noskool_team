package tracks

import (
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
)


type Usecase interface {
	GetTrackByID(trackID int) (*models.Track, error)
	GetTracksByTittle(trackTittle string) ([]*models.Track, error)
	GetTrackByMusicianID(musicianID int) ([]*models.Track, error)
	UploadPicture(trackID int, audioPath string) error
	UploadAudio(trackID int, audioPath string) error
	GetTracksByUserID(userID int) ([]*models.Track, error)
	GetFavoriteTracks(userID int, pagination *commonModels.Pagination) ([]*models.Track, error)
	AddTrackToFavorites(userID, trackID int) error
	DeleteTrackFromFavorites(trackID, userID int) error
	GetTracksByAlbumID(albumID int) ([]*models.Track, error)
	GetTracksByGenreID(genreID int) ([]*models.Track, error)
}
