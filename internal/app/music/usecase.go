package music

import "2021_1_Noskool_team/internal/app/music/models"

type Usecase interface{
	GetMusiciansByGenres(genre string) (*[]models.Musician, error)
}
