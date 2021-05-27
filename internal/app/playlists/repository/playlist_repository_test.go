package repository

import (
	"2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
		"playlist_id", "tittle", "description", "picture", "release_date", "user_id", "uid",
	}).AddRow(playlistsForTest[2].PlaylistID, playlistsForTest[2].Tittle, playlistsForTest[2].Description,
		playlistsForTest[2].Picture, playlistsForTest[2].ReleaseDate, playlistsForTest[2].UserID,
		playlistsForTest[2].UID)

	query := "SELECT playlist_id, tittle, description, picture, release_date, " +
		"user_id, uid FROM playlists\n\t\t\t\t\t\tWHERE playlist_id ="

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	playlist, err := playlistRep.GetPlaylistByID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(playlist, playlistsForTest[2]) {
		t.Fatalf("Not equal")
	}
}

func TestCreatePlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"playlist_id",
	}).AddRow(playlistsForTest[2].PlaylistID)

	queryFirst := "INSERT INTO playlists"

	mock.ExpectQuery(queryFirst).WithArgs(playlistsForTest[2].Tittle, playlistsForTest[2].Description,
		"/api/v1/music/data/img/playlists/happy.webp", playlistsForTest[2].UserID,
	).WillReturnRows(rows)

	updateSecond := "update playlists set uid ="
	mock.ExpectExec(updateSecond).WillReturnResult(sqlmock.NewResult(1, 1))

	querySecond := "INSERT INTO Tracks_to_Playlist"
	mock.ExpectExec(querySecond).WillReturnResult(sqlmock.NewResult(1, 1))

	thirdSecond := "INSERT INTO playlists_to_user"
	mock.ExpectExec(thirdSecond).WithArgs(playlistsForTest[2].UserID, playlistsForTest[2].PlaylistID).WillReturnResult(sqlmock.NewResult(1, 1))

	playlist := &models.Playlist{
		PlaylistID:  3,
		Tittle:      "Tittle without tracks",
		Description: "some description",
		Picture:     "/api/v1/data/img/playlists/happy.webp",
		ReleaseDate: "2020-03-04",
		UserID:      1,
		Tracks: []*trackModels.Track{
			{
				TrackID:     1,
				Tittle:      "song",
				Text:        "sing song song",
				Audio:       "audio",
				Picture:     "picture",
				ReleaseDate: "date",
			},
		},
	}
	playlistAfterCreation, err := playlistRep.CreatePlaylist(playlist)

	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistAfterCreation, playlist) {
		t.Fatalf("Not equal")
	}
}

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

//func TestAddPlaylistToMediateka(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("cant create mock '%s'", err)
//	}
//	playlistRep := NewPlaylistRepository(db)
//	defer db.Close()
//
//	rows := sqlmock.NewRows([]string{
//		"playlist_id", "tittle", "description", "picture", "release_date", "user_id",
//	}).AddRow(playlistsForTest[2].PlaylistID, playlistsForTest[2].Tittle, playlistsForTest[2].Description,
//		playlistsForTest[2].Picture, playlistsForTest[2].ReleaseDate, playlistsForTest[2].UserID)
//
//	queryGetPlaylistID := `SELECT playlist_id, tittle, description, picture, release_date, user_id FROM playlists`
//	mock.ExpectQuery(queryGetPlaylistID).WithArgs(playlistsForTest[2].PlaylistID).WillReturnRows(rows)
//	query := "INSERT"
//	mock.ExpectExec(query).WithArgs(1, 3).WillReturnResult(
//		sqlmock.NewResult(1, 1))
//
//	err = playlistRep.AddPlaylistToMediateka(1, 3)
//
//	assert.NoError(t, err)
//}

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
	query := "select p.playlist_id, p.tittle, p.description, p.picture," +
		"\n       \t\tp.release_date, p.user_id from  playlists p \n\t\t\twhere p.user_id ="

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	playlists, err := playlistRep.GetMediateka(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistsForTest, playlists) {
		t.Fatalf("Not equal")
	}
}

func TestSearchTracks(t *testing.T) {
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
	query := "SELECT p.playlist_id, p.tittle, p.description, p.picture" +
		",\n    \t\tp.release_date, p.user_id FROM playlists as p\n\t\t\tWHERE p.tittle LIKE"

	mock.ExpectQuery(query).WithArgs("some query").WillReturnRows(rows)
	playlists, err := playlistRep.SearchPlaylists("some query")
	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistsForTest, playlists) {
		t.Fatalf("Not equal")
	}
}

func TestGetPlaylistsByGenreID(t *testing.T) {
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
	query := "select p.playlist_id, p.tittle, p.description, p.picture, p.release_date," +
		"\n     \t\tp.user_id from playlists_to_genres as p_g\n\t\t\tleft join playlists " +
		"p on p.playlist_id = p_g.playlist_id\n\t\t\twhere p_g.genre_id"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	playlists, err := playlistRep.GetPlaylistsByGenreID(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistsForTest, playlists) {
		t.Fatalf("Not equal")
	}
}

func TestUploadPicture(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()
	query := "UPDATE playlists SET picture = "
	mock.ExpectExec(query).WithArgs("picture path", 1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = playlistRep.UploadPicture(1, "picture path")

	assert.NoError(t, err)
}

func TestDeletePlaylistFromUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()
	query := "DELETE FROM playlists where playlist_id ="
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = playlistRep.DeletePlaylistFromUser(2, 1)

	assert.NoError(t, err)
}

func TestGetPlaylistsa(t *testing.T) {
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
	query := "select playlist_id, tittle, description, picture, release_date, " +
		"user_id\n\t\t\tfrom playlists\n\t\t\torder by rating"

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
	playlists, err := playlistRep.GetPlaylists()
	assert.NoError(t, err)
	if !reflect.DeepEqual(playlistsForTest, playlists) {
		t.Fatalf("Not equal")
	}
}

func TestAddTrackToPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()
	query := `INSERT INTO tracks_to_playlist`
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = playlistRep.AddTrackToPlaylist(2, 1)

	assert.NoError(t, err)
}

func TestDeleteTrackFromPlaylist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	playlistRep := NewPlaylistRepository(db)

	defer db.Close()
	query := `DELETE FROM tracks_to_playlist
			WHERE track_id = `
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = playlistRep.DeleteTrackFromPlaylist(2, 1)

	assert.NoError(t, err)
}
