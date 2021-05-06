package usecase

import (
	mock_playlists "2021_1_Noskool_team/internal/app/playlists/mocks"
	"2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	playlistsForTest = []*models.Playlist{
		{
			PlaylistID:  1,
			Tittle:      "Tittle of first playlist",
			Description: "some description",
			Picture:     "/api/v1/data/img/playlists/1.png",
			ReleaseDate: "2020-03-04",
			UserID:      1,
		},
		{
			PlaylistID:  2,
			Tittle:      "Tittle of second playlist",
			Description: "some other description",
			Picture:     "/api/v1/data/img/playlists/1.png",
			ReleaseDate: "2020-07-07",
			UserID:      5,
		},
		{
			PlaylistID:  3,
			Tittle:      "Tittle without tracks",
			Description: "some description",
			Picture:     "/api/v1/data/img/playlists/3.png",
			ReleaseDate: "2020-03-04",
			UserID:      1,
		},
	}
	tracksForTest = []*trackModels.Track{
		{
			TrackID:     1,
			Tittle:      "song",
			Text:        "sing song song",
			Audio:       "audio",
			Picture:     "picture",
			ReleaseDate: "date",
		},
		{
			TrackID:     2,
			Tittle:      "dsfds",
			Text:        "sifdsfdsg song song",
			Audio:       "afdsudio",
			Picture:     "fdsafdsa",
			ReleaseDate: "dafdste",
		},
	}
)

func TestGetPlaylistByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playlistRepMock := mock_playlists.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepMock)

	playlistRepMock.
		EXPECT().GetPlaylistByID(gomock.Eq(1)).
		Return(playlistsForTest[0], nil)
	playlistRepMock.
		EXPECT().GetTracksByPlaylistID(gomock.Eq(1)).
		Return(tracksForTest, nil)

	playlistsForTest[0].Tracks = tracksForTest

	result, err := playlistUsecase.GetPlaylistByID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, result, playlistsForTest[0])
}

func TestGetMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	playlistRepMock := mock_playlists.NewMockRepository(ctrl)
	playlistUsecase := NewPlaylistUsecase(playlistRepMock)

	playlistRepMock.
		EXPECT().GetMediateka(gomock.Eq(1)).AnyTimes().
		Return(playlistsForTest, nil)
	playlistRepMock.
		EXPECT().GetTracksByPlaylistID(gomock.Any()).AnyTimes().
		Return(tracksForTest, nil)

	for _, playlist := range playlistsForTest {
		playlist.Tracks = tracksForTest
	}

	result, err := playlistUsecase.GetMediateka(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, result, playlistsForTest)
}

func TestGetPlaylists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().GetPlaylists().Return(playlistsForTest, nil)
	mockRepo.EXPECT().GetTracksByPlaylistID(gomock.Any()).AnyTimes()
	playlists, err := mockUsecase.GetPlaylists()
	assert.Equal(t, err, nil)
	assert.Equal(t, playlistsForTest, playlists)
}

func TestCreatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().CreatePlaylist(playlistsForTest[0]).Return(playlistsForTest[0], nil)
	playlists, err := mockUsecase.CreatePlaylist(playlistsForTest[0])
	assert.Equal(t, err, nil)
	assert.Equal(t, playlistsForTest[0], playlists)
}

func TestDeletePlaylistFromUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().DeletePlaylistFromUser(1, 2).Return(nil)
	err := mockUsecase.DeletePlaylistFromUser(1, 2)
	assert.Equal(t, err, nil)
}

func TestAddPlaylistToMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().AddPlaylistToMediateka(1, 2).Return(nil)
	err := mockUsecase.AddPlaylistToMediateka(1, 2)
	assert.Equal(t, err, nil)
}

func TestUploadAudio(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().UploadPicture(1, "some path").Return(nil)
	err := mockUsecase.UploadPicture(1, "some path")
	assert.Equal(t, err, nil)
}

func TestGetPlaylistsByGenreID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_playlists.NewMockRepository(ctrl)
	mockUsecase := NewPlaylistUsecase(mockRepo)

	mockRepo.EXPECT().GetPlaylistsByGenreID(1).Return(playlistsForTest, nil)
	playlists, err := mockUsecase.GetPlaylistsByGenreID(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, playlistsForTest, playlists)
}
