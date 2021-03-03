package usecase

import (
	mocktracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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

	track, err := mockUsecase.GetTrackByMusicianID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedTrack, track)
}