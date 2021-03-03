package usecase

import (
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	_ "github.com/lib/pq" //goland:noinspection
)

type TracksUsecase struct {
	trackRep tracks.Repository
}

func NewTracksUsecase(trackRep tracks.Repository) *TracksUsecase {
	return &TracksUsecase{
		trackRep: trackRep,
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
