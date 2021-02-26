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

func NewMusicsUsecase(config *configs.Config) MusicUsecase {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 dbname=music_service sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
	}
	return MusicUsecase{
		musicRepo: repository.NewMusicRepository(db),
	}
}

func (usecase *MusicUsecase) GetMusiciansByGenres(genre string) (*[]models.Musician, error) {
	musicians, err := usecase.musicRepo.GetMusiciansByGenres(genre)
	if err != nil {
		return nil, err
	}
	return &musicians, nil
}

func (usecase *MusicUsecase) GetMusic() {

}
