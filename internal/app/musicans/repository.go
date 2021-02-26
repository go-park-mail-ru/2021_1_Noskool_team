package musicans

import (
	"2021_1_Noskool_team/internal/app/musicans/models"
)

type Repository interface {
	GetTrackById(trackId int) (models.Track, error)
	GetMusiciansByGenres(genre string) ([]models.Musician, error)
}
