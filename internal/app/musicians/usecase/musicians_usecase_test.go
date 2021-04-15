package usecase

import (
	mockMusicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
			Name:        "Кровосток",
			Description: "Российская рэп-группа из Москвы",
			Picture:     "picture of Кровосток",
		},
	}

	mockRepo.
		EXPECT().GetMusiciansByGenre(gomock.Eq("rok")).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusiciansByGenre("rok")
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}

func TestGetMusicianByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &[]models.Musician{
		{
			MusicianID:  1,
			Name:        "Oliver Tree",
			Description: "Американский певец, продюсер и режиссёр из города Санта-Круз",
			Picture:     "picture of Oliver Tree",
		},
	}

	mockRepo.
		EXPECT().GetMusicianByTrackID(gomock.Eq(1)).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusicianByTrackID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}

func TestGetMusicianByAlbumID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &[]models.Musician{
		{
			MusicianID:  1,
			Name:        "Увула",
			Description: "Группа Увула родилась в Петербурге несколько лет назад",
			Picture:     "picture of Увула",
		},
	}

	mockRepo.
		EXPECT().GetMusicianByAlbumID(gomock.Eq(1)).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusicianByAlbumID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}

func TestGetMusicianByPlaylistID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &[]models.Musician{
		{
			MusicianID:  1,
			Name:        "Пасош",
			Description: "Российская рок-группа из Москвы",
			Picture:     "picture of Пасош",
		},
	}

	mockRepo.
		EXPECT().GetMusicianByPlaylistID(gomock.Eq(1)).
		Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusicianByPlaylistID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}

func TestGetMusiciansTop3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockMusicians.NewMockRepository(ctrl)
	mockUsecase := NewMusicsUsecase(mockRepo)

	expectedMusician := &[]models.Musician{
		{
			MusicianID:  1,
			Name:        "Пасош",
			Description: "Российская рок-группа из Москвы",
			Picture:     "picture of Пасош",
		},
	}

	mockRepo.EXPECT().GetMusiciansTop4().Return(expectedMusician, nil)

	musician, err := mockUsecase.GetMusiciansTop4()
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedMusician, musician)
}
