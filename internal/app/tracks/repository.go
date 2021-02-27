package tracks

import "2021_1_Noskool_team/internal/app/tracks/models"

type Repository interface {
	GetTrackByID(trackID int) (*models.Track, error)
	GetTracksByTittle(trackTittle string) ([]*models.Track, error)
	GetTrackByMusicianID(musicianID int) ([]*models.Track, error)
}
