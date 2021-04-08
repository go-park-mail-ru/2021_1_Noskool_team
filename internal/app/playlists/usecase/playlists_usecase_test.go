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
