package usecase

import (
	mocktracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tracksForTests = []*models.Track{
		{
			TrackID:     1,
			Tittle:      "song",
			Text:        "sing song song",
			Audio:       "/api/v1/data/audio/track/2.mp3",
			Picture:     "picture",
			ReleaseDate: "2021-03-04",
		},
		{
			TrackID:     2,
			Tittle:      "song helloWorld",
			Text:        "sing song song ooooo",
			Audio:       "/api/v1/data/audio/2.mp3",
			Picture:     "/api/v1/data/audio/tracks/2.mp3",
			ReleaseDate: "2020-03-04",
		},
	}
)

func TestGetTrackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	expectedTrack := &models.Track{
		TrackID:     1,
		Tittle:      "tittle",
		Text:        "text",
		Audio:       "audio",
		Picture:     "picture",
		ReleaseDate: "date",
	}

	mockRepo.
		EXPECT().GetTrackByID(gomock.Eq(expectedTrack.TrackID)).
		Return(expectedTrack, nil)
	mockRepo.EXPECT().GetMusicianByTrackID(gomock.Eq(expectedTrack.TrackID))
	mockRepo.EXPECT().GetGenreByTrackID(gomock.Eq(expectedTrack.TrackID))
	mockRepo.EXPECT().GetAlbumsByTrackID(gomock.Eq(expectedTrack.TrackID))

	track, err := mockUsecase.GetTrackByID(expectedTrack.TrackID)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedTrack, track)
}

func TestGetTracksByTittle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	expectedTrack := []*models.Track{
		{
			TrackID:     1,
			Tittle:      "tittle",
			Text:        "text",
			Audio:       "audio",
			Picture:     "picture",
			ReleaseDate: "date",
		},
	}

	mockRepo.
		EXPECT().GetTracksByTittle(gomock.Eq(expectedTrack[0].Tittle)).
		Return(expectedTrack, nil)

	track, err := mockUsecase.GetTracksByTittle(expectedTrack[0].Tittle)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedTrack, track)
}

func TestGetTrackByMusicianID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	expectedTrack := []*models.Track{
		{
			TrackID:     1,
			Tittle:      "tittle",
			Text:        "text",
			Audio:       "audio",
			Picture:     "picture",
			ReleaseDate: "date",
		},
	}

	mockRepo.
		EXPECT().GetTrackByMusicianID(gomock.Eq(1)).
		Return(expectedTrack, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(expectedTrack).Return(expectedTrack)

	track, err := mockUsecase.GetTrackByMusicianID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedTrack, track)
}

func TestAddDeleteTrackToMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().DeleteTrackFromMediateka(gomock.Eq(1), gomock.Eq(1)).
		Return(nil)
	err := mockUsecase.AddDeleteTrackToMediateka(1, 1, "delete")
	assert.Equal(t, err, nil)

	mockRepo.EXPECT().AddTrackToMediateka(gomock.Eq(1), gomock.Eq(1)).
		Return(nil)
	err = mockUsecase.AddDeleteTrackToMediateka(1, 1, "add")
	assert.Equal(t, err, nil)

	err = mockUsecase.AddDeleteTrackToMediateka(1, 1, "some error type")
	assert.Equal(t, err, errors.New("unknown operation"))
}

func TestGetTracksByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetTracksByUserID(gomock.Eq(1)).
		Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetTracksByUserID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestGetTracksByAlbumID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetTracksByAlbumID(gomock.Eq(1)).
		Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetTracksByAlbumID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestGetTracksByGenreID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetTracksByGenreID(gomock.Eq(1)).
		Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetTracksByGenreID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestGetTop20Tracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetTop20Tracks().
		Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetTop20Tracks()
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestGetBillbordTopCharts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetBillbordTopCharts().Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetBillbordTopCharts()
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestGetHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetHistory(1).
		Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(tracksForTests).Return(tracksForTests)
	track, err := mockUsecase.GetHistory(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestSearchTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().SearchTracks("search query").
		Return(tracksForTests, nil)
	track, err := mockUsecase.SearchTracks("search query")
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestUploadPicture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().UploadPicture(1, "some path").
		Return(nil)
	err := mockUsecase.UploadPicture(1, "some path")
	assert.Equal(t, err, nil)
}

func TestUploadAudio(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().UploadAudio(1, "some path").
		Return(nil)
	err := mockUsecase.UploadAudio(1, "some path")
	assert.Equal(t, err, nil)
}

func TestAddTrackToFavorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().CheckTrackInFavorite(1, 2).
		Return(fmt.Errorf("some error"))
	mockRepo.EXPECT().IncrementLikes(2).
		Return(nil)
	mockRepo.EXPECT().CheckTrackInMediateka(1, 2).
		Return(nil)
	mockRepo.EXPECT().AddTrackToFavorites(1, 2).
		Return(nil)
	err := mockUsecase.AddTrackToFavorites(1, 2)
	assert.Equal(t, err, nil)
}

func TestDeleteTrackFromFavorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().CheckTrackInFavorite(1, 2).
		Return(nil)
	mockRepo.EXPECT().DecrementLikes(2).
		Return(nil)
	mockRepo.EXPECT().DeleteTrackFromFavorites(1, 2).
		Return(nil)
	err := mockUsecase.DeleteTrackFromFavorites(1, 2)
	assert.Equal(t, err, nil)
}

func TestAddToHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().AddToHistory(1, 2).
		Return(nil)
	err := mockUsecase.AddToHistory(1, 2)
	assert.Equal(t, err, nil)
}

func TestGetTopTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	mockRepo.EXPECT().GetTopTrack().Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(gomock.Any()).Return(tracksForTests)
	track, err := mockUsecase.GetTopTrack()
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}

func TestCheckTrackInMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)
	mockRepo.EXPECT().CheckTrackInMediateka(1, 2).Return(nil)

	result := mockUsecase.CheckTrackInMediateka(1, 2)
	assert.Equal(t, true, result)
}

func TestCheckTrackInFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)
	mockRepo.EXPECT().CheckTrackInFavorite(1, 2).Return(nil)

	result := mockUsecase.CheckTrackInFavorite(1, 2)
	assert.Equal(t, true, result)
}

func TestGetFavoriteTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocktracks.NewMockRepository(ctrl)
	mockUsecase := NewTracksUsecase(mockRepo)

	pagination := &commonModels.Pagination{
		Limit:  1,
		Offset: 2,
	}

	mockRepo.EXPECT().GetFavoriteTracks(1, pagination).Return(tracksForTests, nil)
	mockRepo.EXPECT().GetMusiciansGenresAndAlbums(gomock.Any()).Return(tracksForTests)
	track, err := mockUsecase.GetFavoriteTracks(1, pagination)
	assert.Equal(t, err, nil)
	assert.Equal(t, tracksForTests, track)
}
