package usecase

import (
	mockAlbum "2021_1_Noskool_team/internal/app/album/mocks"
	albumModels "2021_1_Noskool_team/internal/app/album/models"
	mockMusicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	musicianModels "2021_1_Noskool_team/internal/app/musicians/models"
	mockPlaylists "2021_1_Noskool_team/internal/app/playlists/mocks"
	playlistModels "2021_1_Noskool_team/internal/app/playlists/models"
	"2021_1_Noskool_team/internal/app/search/models"
	mockTracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	models2 "2021_1_Noskool_team/internal/app/tracks/models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	tr = []*models2.Track{
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
	searchResultForTests = &models.Search{
		Tracks: []*models.TrackWithAlbum{
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
		},
		Albums: []*albumModels.Album{
			{
				AlbumID:     1,
				Tittle:      "album1",
				Picture:     "picture1",
				ReleaseDate: "date1",
			},
			{
				AlbumID:     2,
				Tittle:      "album2",
				Picture:     "picture2",
				ReleaseDate: "date2",
			},
		},
		Musicians: []*musicianModels.Musician{
			{
				MusicianID:  1,
				Name:        "Joji",
				Description: "Pretty Boy",
				Picture:     "picture",
			},
			{
				MusicianID:  2,
				Name:        "Дора",
				Description: "Дура",
				Picture:     "picture",
			},
		},
		Playlists: []*playlistModels.Playlist{
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
		},
	}

	noContentResult = &models.Search{
		Tracks:    []*models.TrackWithAlbum{},
		Albums:    []*albumModels.Album{},
		Musicians: []*musicianModels.Musician{},
		Playlists: []*playlistModels.Playlist{},
	}
)

func TestSearchContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trackRepMock := mockTracks.NewMockRepository(ctrl)
	playlistRepMock := mockPlaylists.NewMockRepository(ctrl)
	albumRepMock := mockAlbum.NewMockRepository(ctrl)
	musicianRepMock := mockMusicians.NewMockRepository(ctrl)

	searchUsecase := NewSearchUsecase(trackRepMock, albumRepMock,
		musicianRepMock, playlistRepMock)

	searchQuery := "some query"

	trackRepMock.
		EXPECT().SearchTracks(gomock.Eq(searchQuery)).
		Return(tr, nil)
	albumRepMock.
		EXPECT().GetAlbumsByTrackID(gomock.Any()).AnyTimes()
	playlistRepMock.
		EXPECT().SearchPlaylists(gomock.Eq(searchQuery)).
		Return(searchResultForTests.Playlists, nil)
	albumRepMock.
		EXPECT().SearchAlbums(gomock.Eq(searchQuery)).
		Return(searchResultForTests.Albums, nil)
	musicianRepMock.
		EXPECT().SearchMusicians(gomock.Eq(searchQuery)).
		Return(searchResultForTests.Musicians, nil)

	searchResult := searchUsecase.SearchContent(searchQuery)
	assert.Equal(t, searchResult, searchResultForTests)

	returnErr := errors.New("no content")

	trackRepMock.
		EXPECT().SearchTracks(gomock.Eq(searchQuery)).
		Return(nil, returnErr)
	playlistRepMock.
		EXPECT().SearchPlaylists(gomock.Eq(searchQuery)).
		Return(nil, returnErr)
	albumRepMock.
		EXPECT().SearchAlbums(gomock.Eq(searchQuery)).
		Return(nil, returnErr)
	musicianRepMock.
		EXPECT().SearchMusicians(gomock.Eq(searchQuery)).
		Return(nil, returnErr)

	searchResult = searchUsecase.SearchContent(searchQuery)
	assert.Equal(t, searchResult, noContentResult)
}
