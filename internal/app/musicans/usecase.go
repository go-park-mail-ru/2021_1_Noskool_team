package musicans

import "2021_1_Noskool_team/internal/app/musicans/models"

type Usecase interface {
	GetMusiciansByGenres(genre string) (*[]models.Musician, error)
}
