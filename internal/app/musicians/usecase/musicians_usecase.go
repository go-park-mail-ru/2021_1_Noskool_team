package usecase

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/musicians/repository"
	"database/sql"
	"fmt"
)

type MusicUsecase struct {
	musicRepo musicians.Repository
}

func NewMusicsUsecase(config *configs.Config) *MusicUsecase {
	db, err := sql.Open("postgres",
		config.MusicPostgresBD,
	)
	if err != nil {
		fmt.Println(err)
	}
	return &MusicUsecase{
		musicRepo: repository.NewMusicRepository(db),
	}
}

func (usecase *MusicUsecase) GetMusiciansByGenres(genre string) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusiciansByGenres(genre)
	if err != nil {
		return nil, err
	}
	return &mus, nil
}

func (usecase *MusicUsecase) GetMusic() {

}

func (usecase *MusicUsecase) GetMusicianByID(musicianID int) (*models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicianByID(musicianID)
	if err != nil {
		return nil, err
	}
	return mus, nil
}
