package repository

import (
	"2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestGetTrackByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"playlist_id", "tittle", "description", "picture", "release_date", "user_id",
	}).AddRow(playlistsForTest[2].PlaylistID, playlistsForTest[2].Tittle, playlistsForTest[2].Description,
		playlistsForTest[2].Picture, playlistsForTest[2].ReleaseDate, playlistsForTest[2].UserID)

	query := "SELECT playlist_id, tittle, description, picture, release_date, " +
		"user_id FROM playlists\n\t\t\t\t\t\tWHERE playlist_id ="

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	playlist, err := playlistRep.GetPlaylistByID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(playlist, playlistsForTest[2]) {
		t.Fatalf("Not equal")
	}
}

//func TestCreatePlaylist(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("cant create mock '%s'", err)
//	}
//	playlistRep := NewPlaylistRepository(db)
//
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{
//		"playlist_id",
//	}).AddRow(playlistsForTest[2].PlaylistID)
//
//	queryFirst := "INSERT INTO playlists"
//
//	mock.ExpectQuery(queryFirst).WithArgs(playlistsForTest[2].Tittle, playlistsForTest[2].Description,
//		playlistsForTest[2].Picture, playlistsForTest[2].ReleaseDate, playlistsForTest[2].UserID,
//	).WillReturnRows(rows)
//
//	querySecond := "INSERT INTO Tracks_to_Playlist"
//	mock.ExpectExec(querySecond).WithArgs(playlistsForTest[2].UserID, playlistsForTest[2].PlaylistID,
//	).WillReturnResult(sqlmock.NewResult(1, 1))
//
//	thirdSecond := "INSERT INTO playlists_to_user"
//	mock.ExpectExec(thirdSecond).WithArgs(playlistsForTest[2].UserID, playlistsForTest[2].PlaylistID,
//	).WillReturnResult(sqlmock.NewResult(1, 1))
//
//	playlist, err := playlistRep.CreatePlaylist(playlistsForTest[2])
//
//	assert.NoError(t, err)
//	if !reflect.DeepEqual(playlist, playlistsForTest[2]) {
//		t.Fatalf("Not equal")
//	}
//}

func TestGetTracksByPlaylistID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	})
	for _, row := range tracksForTest {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate)
	}

	query := "select t.track_id, t.tittle, t.text, t.audio, t.picture, t.release_date " +
		"from tracks_to_playlist as t_p\n\t\t\tleft outer join tracks as t on t.track_id = t_p.track_id"
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	track, err := playlistRep.GetTracksByPlaylistID(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTest, track) {
		t.Fatalf("Not equal")
	}
}

func TestAddPlaylistToMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()
	query := "INSERT INTO playlists_to_user"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = playlistRep.AddPlaylistToMediateka(1, 2)

	assert.NoError(t, err)
}

func TestGetMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"playlist_id", "tittle", "description", "picture", "release_date", "user_id",
	})
	for _, row := range playlistsForTest {
		rows.AddRow(row.PlaylistID, row.Tittle, row.Description,
			row.Picture, row.ReleaseDate, row.UserID)
	}
	query := "select p.playlist_id, p.tittle, p.description, p.picture,\n       " +
		"p.release_date, p.user_id from playlists_to_user " +
		"as p_u\n\t\t\tleft join playlists p on p_u.playlist_id = p.playlist_id"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	playlists, err := playlistRep.GetMediateka(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistsForTest, playlists) {
		t.Fatalf("Not equal")
	}
}
