package usecase

import (
	"2021_1_Noskool_team/configs"
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	"2021_1_Noskool_team/internal/app/tracks/repository"
	"database/sql"
	"fmt"
)

type TracksUsecase struct {
	trackRep tracks.Repository
}

func NewTracksUsecase(config *configs.Config) TracksUsecase {
	dbCon, err := sql.Open("postgres",
		config.MusicPostgresBD,
	)
	if err != nil {
		fmt.Println(err)
	}
	return TracksUsecase{
		trackRep: repository.NewTracksRepository(dbCon),
	}
}

func (usecase *TracksUsecase) GetTrackByID(trackID int) (*models.Track, error) {
	track, err := usecase.trackRep.GetTrackByID(trackID)
	return track, err
}

func (usecase *TracksUsecase) GetTracksByTittle(trackTittle string) ([]*models.Track, error) {
	track, err := usecase.trackRep.GetTracksByTittle(trackTittle)
	return track, err
}

func (usecase *TracksUsecase) GetTrackByMusicianID(musicianID int) ([]*models.Track, error) {
	track, err := usecase.trackRep.GetTrackByMusicianID(musicianID)
	return track, err
}