package tracks

import "2021_1_Noskool_team/internal/app/tracks/models"

type Repository interface {
	GetTrackByID(trackID int) (*models.Track, error)
	GetTracksByTittle(trackTittle string) ([]*models.Track, error)
	GetTrackByMusicianID(musicianID int) ([]*models.Track, error)
	CreateTrack(*models.Track) (*models.Track, error)
	UploadPicture(trackID int, audioPath string) error
	UploadAudio(trackID int, audioPath string) error
	GetTracksByUserID(userID int) ([]*models.Track, error)
	GetFavoriteTracks(userID int) ([]*models.Track, error)
}
