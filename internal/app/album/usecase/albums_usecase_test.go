package usecase

import (
	mockalbums "2021_1_Noskool_team/internal/app/album/mocks"
	"2021_1_Noskool_team/internal/app/album/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockalbums.NewMockRepository(ctrl)
	mockUsecase := NewAlbumcUsecase(mockRepo)

	expectedAlbum := &models.Album{
		AlbumID:     1,
		Tittle:      "Some album",
		Picture:     "picture of album",
		ReleaseDate: "zavtra",
	}

	mockRepo.EXPECT().GetAlbumByID(gomock.Eq(expectedAlbum.AlbumID)).Return(expectedAlbum, nil)
	track, err := mockUsecase.GetAlbumByID(expectedAlbum.AlbumID)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedAlbum, track)
}

func TestGetAlbumsByMusicianID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockalbums.NewMockRepository(ctrl)
	mockUsecase := NewAlbumcUsecase(mockRepo)

	expectedAlbum := &[]models.Album{
		{
			AlbumID:     1,
			Tittle:      "Some album",
			Picture:     "picture of album",
			ReleaseDate: "zavtra",
		},
	}

	mockRepo.EXPECT().GetAlbumsByMusicianID(gomock.Eq(1)).Return(expectedAlbum, nil)
	track, err := mockUsecase.GetAlbumsByMusicianID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedAlbum, track)
}

func TestGetAlbumsByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockalbums.NewMockRepository(ctrl)
	mockUsecase := NewAlbumcUsecase(mockRepo)

	expectedAlbum := &[]models.Album{
		{
			AlbumID:     1,
			Tittle:      "Some album",
			Picture:     "picture of album",
			ReleaseDate: "zavtra",
		},
	}

	mockRepo.EXPECT().GetAlbumsByTrackID(gomock.Eq(1)).Return(expectedAlbum, nil)
	track, err := mockUsecase.GetAlbumsByTrackID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedAlbum, track)
}

func TestSearchAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mockalbums.NewMockRepository(ctrl)
	mockUsecase := NewAlbumcUsecase(mockRepo)

	expectedAlbum := []*models.Album{
		{
			AlbumID:     1,
			Tittle:      "Some album",
			Picture:     "picture of album",
			ReleaseDate: "zavtra",
		},
	}

	mockRepo.EXPECT().SearchAlbums(gomock.Eq("1")).Return(expectedAlbum, nil)
	track, err := mockUsecase.SearchAlbums("1")
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedAlbum, track)
}
