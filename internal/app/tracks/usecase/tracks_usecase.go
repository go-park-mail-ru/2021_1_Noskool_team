package usecase

import (
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"errors"
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

func (usecase *TracksUsecase) UploadPicture(trackID int, audioPath string) error {
	err := usecase.trackRep.UploadPicture(trackID, audioPath)
	return err
}
func (usecase *TracksUsecase) UploadAudio(trackID int, audioPath string) error {
	err := usecase.trackRep.UploadAudio(trackID, audioPath)
	return err
}

func (usecase *TracksUsecase) GetTracksByUserID(userID int) ([]*models.Track, error) {
	tracks, err := usecase.trackRep.GetTracksByUserID(userID)
	return tracks, err
}

func (usecase *TracksUsecase) GetFavoriteTracks(userID int,
	pagination *commonModels.Pagination) ([]*models.Track, error) {
	tracks, err := usecase.trackRep.GetFavoriteTracks(userID, pagination)
	return tracks, err
}

func (usecase *TracksUsecase) AddTrackToFavorites(userID, trackID int) error {
	err := usecase.trackRep.AddTrackToFavorites(userID, trackID)
	return err
}

func (usecase *TracksUsecase) DeleteTrackFromFavorites(userID, trackID int) error {
	err := usecase.trackRep.DeleteTrackFromFavorites(userID, trackID)
	return err
}

func (usecase *TracksUsecase) GetTracksByAlbumID(albumID int) ([]*models.Track, error) {
	tracksByAlbum, err := usecase.trackRep.GetTracksByAlbumID(albumID)
	return tracksByAlbum, err
}

func (usecase *TracksUsecase) GetTracksByGenreID(genreID int) ([]*models.Track, error) {
	tracksByGenre, err := usecase.trackRep.GetTracksByGenreID(genreID)
	return tracksByGenre, err
}

func (usecase *TracksUsecase) AddDeleteTrackToMediateka(userID, trackID int, operationType string) error {
	var err error
	if operationType == "add" {
		err = usecase.trackRep.AddTrackToMediateka(userID, trackID)
	} else if operationType == "delete" {
		err = usecase.trackRep.DeleteTrackFromMediateka(userID, trackID)
	} else {
		return errors.New("unknown operation")
	}
	return err
}

func (usecase *TracksUsecase) SearchTracks(searchQuery string) ([]*models.Track, error) {
	tracks, err := usecase.trackRep.SearchTracks(searchQuery)
	return tracks, err
}

func (usecase *TracksUsecase) GetTop20Tracks() ([]*models.Track, error) {
	tracks, err := usecase.trackRep.GetTop20Tracks()
	return tracks, err
}

func (usecase *TracksUsecase) GetBillbordTopCharts() ([]*models.Track, error) {
	tracks, err := usecase.trackRep.GetTop20Tracks()
	return tracks, err
}
