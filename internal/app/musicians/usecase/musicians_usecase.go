package usecase

import (
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"

	_ "github.com/lib/pq" //goland:noinspection
)

type MusicUsecase struct {
	musicRepo musicians.Repository
}

func NewMusicsUsecase(musicRep musicians.Repository) *MusicUsecase {
	return &MusicUsecase{
		musicRepo: musicRep,
	}
}

func (usecase *MusicUsecase) GetMusiciansByGenres(genre string) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusiciansByGenres(genre)
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetMusicianByID(musicianID int) (*models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicianByID(musicianID)
	if err != nil {
		return nil, err
	}
	return mus, nil
}
