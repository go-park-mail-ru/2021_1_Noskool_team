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

func (usecase *MusicUsecase) GetMusiciansTop4() (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusiciansTop4()
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetMusicians() (*[]models.Musician, error) {
	mus, err := usecase.musicRepo.GetMusicians()
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) GetGenreForMusician(nameMusician string) (*[]string, error) {
	mus, err := usecase.musicRepo.GetGenreForMusician(nameMusician)
	if err != nil {
		return nil, err
	}
	return mus, nil
}

func (usecase *MusicUsecase) AddMusicianToMediateka(userID, musicianID int) error {
	err := usecase.musicRepo.AddMusicianToMediateka(userID, musicianID)
	return err
}

func (usecase *MusicUsecase) CheckMusicianInFavorite(userID, musicianID int) error {
	return usecase.musicRepo.CheckMusicianInFavorite(userID, musicianID)
}

func (usecase *MusicUsecase) CheckMusicianInMediateka(userID, musicianID int) error {
	return usecase.musicRepo.CheckMusicianInMediateka(userID, musicianID)
}

func (usecase *MusicUsecase) AddMusicianToFavorites(userID, musicianID int) error {
	err := usecase.musicRepo.CheckMusicianInMediateka(userID, musicianID)
	if err != nil {
		err = usecase.musicRepo.AddMusicianToMediateka(userID, musicianID)
		if err != nil {
			return err
		}
	}

	err = usecase.musicRepo.AddMusicianToFavorites(userID, musicianID)
	return err
}

func (usecase *MusicUsecase) DeleteMusicianFromFavorites(userID, musicianID int) error {
	err := usecase.musicRepo.DeleteMusicianFromFavorites(userID, musicianID)
	return err
}

func (usecase *MusicUsecase) GetMusiciansMediateka(userID int) ([]*models.Musician, error) {
	musicians, err := usecase.musicRepo.GetMusiciansMediateka(userID)
	return musicians, err
}

func (usecase *MusicUsecase) GetMusiciansFavorites(userID int) ([]*models.Musician, error) {
	musicians, err := usecase.musicRepo.GetMusiciansFavorites(userID)
	return musicians, err
}

func (usecase *MusicUsecase) DeleteMusicianFromMediateka(userID, musicianID int) error {
	err := usecase.musicRepo.DeleteMusicianFromMediateka(userID, musicianID)
	return err
}
