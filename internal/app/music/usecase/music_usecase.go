package usecase

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/music"
)

type MusicUsecase struct {
	musicRepo music.Repository
}

func NewSessionsUsecase(config *configs.Config) MusicUsecase {
	return MusicUsecase{

	}
}

func (usecase *MusicUsecase) GetMusic() {

}
