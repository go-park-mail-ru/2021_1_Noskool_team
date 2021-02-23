package music

import "2021_1_Noskool_team/internal/app/music/models"

type Repository interface {
	GetTrackById(trackId int) (models.Track, error)
}