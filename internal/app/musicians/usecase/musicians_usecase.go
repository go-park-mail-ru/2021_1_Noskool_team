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

func (usecase *MusicUsecase) GetMusiciansByGenre(genre string) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusiciansByGenre(genre)
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

func (usecase *MusicUsecase) GetMusicianByTrackID(trackID int) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicianByTrackID(trackID)
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetMusicianByAlbumID(albumID int) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicianByAlbumID(albumID)
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicianByPlaylistID(playlistID)
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetMusiciansTop3() (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusiciansTop3()
	if err != nil {
		return nil, err
	}
	return mus, nil
}
