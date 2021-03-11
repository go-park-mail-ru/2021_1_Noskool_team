package musicians

import "2021_1_Noskool_team/internal/app/musicians/models"

type Usecase interface {
	GetMusiciansByGenres(genre string) (*[]models.Musician, error)
	GetMusicianByID(musicianID int) (*models.Musician, error)
}
