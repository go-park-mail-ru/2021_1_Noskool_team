package usecase

import (
	mockMusicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMusicianByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &models.Musician{
		MusicianID:  1,
		Name:        "Anton",
		Description: "Anton Krutoy",
		Picture:     "picture of anton",
	}

	mockRepo.
		EXPECT().GetMusicianByID(gomock.Eq(expectedMusician.MusicianID)).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusicianByID(expectedMusician.MusicianID)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}

func TestGetMusiciansByGenres(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &[]models.Musician{
		{
			MusicianID:  1,
			Name:        "Anton",
			Description: "Anton Krutoy",
			Picture:     "picture of anton",
		},
	}

	mockRepo.
		EXPECT().GetMusiciansByGenres(gomock.Eq("rok")).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusiciansByGenres("rok")
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}
