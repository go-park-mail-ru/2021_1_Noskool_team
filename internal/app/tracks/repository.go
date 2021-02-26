package tracks

import "2021_1_Noskool_team/internal/app/tracks/models"

type Repository interface {
	GetTrackById(trackId int) (*models.Track, error)
	GetTracksByTittle(trackTittle string) ([]*models.Track, error)
}
