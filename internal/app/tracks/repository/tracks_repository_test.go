package repository

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

	musiciansForTests = []*musiciansModels.Musician{
		{
			MusicianID:  1,
			Name:        "Some Name",
			Description: "some description",
			Picture:     "some picture",
		},
		{
			MusicianID:  2,
			Name:        "Some Name",
			Description: "some description",
			Picture:     "some picture",
		},
	}

	albumsForTests = []*albumsModels.Album{
		{
			AlbumID:     1,
			Tittle:      "albumalbum1",
			Picture:     "picturepicture1",
			ReleaseDate: "datedate1",
		},
		{
			AlbumID:     2,
			Tittle:      "albumalbum2",
			Picture:     "picturepicture2",
			ReleaseDate: "datedate2",
		},
	}

	genresForTests = []*commonModels.Genre{
		{
			GenreID: 1,
			Title:   "some title",
		},
		{
			GenreID: 2,
			Title:   "some title",
		},
	}
)

func TestGetTrackByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	}).AddRow(1, "song", "sing song song", "audio", "picture", "date", "3:50", 0)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)

	track, err := tracRep.GetTrackByID(1)

	fmt.Println(track)
	assert.NoError(t, err)
	assert.NotNil(t, track)
}

func TestGetTrackByMusicianID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	expectedTrack := &models.Track{
		TrackID:     1,
		Tittle:      "song",
		Text:        "sing song song",
		Audio:       "audio",
		Picture:     "picture",
		ReleaseDate: "date",
		Duration:    "3:50",
	}

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	}).AddRow(expectedTrack.TrackID, expectedTrack.Tittle, expectedTrack.Text,
		expectedTrack.Audio, expectedTrack.Picture, expectedTrack.ReleaseDate,
		expectedTrack.Duration, expectedTrack.Likes)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)

	track, err := tracRep.GetTrackByMusicianID(1)

	fmt.Println(track)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedTrack, track[0]) {
		t.Fatalf("Not equal")
	}
}

func TestGetTracksByTittle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	expectedTrack := &models.Track{
		TrackID:     1,
		Tittle:      "song",
		Text:        "sing song song",
		Audio:       "audio",
		Picture:     "picture",
		ReleaseDate: "date",
		Duration:    "3:50",
	}

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	}).AddRow(expectedTrack.TrackID, expectedTrack.Tittle, expectedTrack.Text,
		expectedTrack.Audio, expectedTrack.Picture, expectedTrack.ReleaseDate,
		expectedTrack.Duration, expectedTrack.Likes)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs("hello").WillReturnRows(rows)

	track, err := tracRep.GetTracksByTittle("hello")

	fmt.Println(track)
	assert.NoError(t, err)
}

func TestGetTrackByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}
	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes from " +
		"tracks\n\t\t\tLEFT JOIN tracks_to_user"
	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)
	track, err := tracRep.GetTracksByUserID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetTrackByAlbumID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}
	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes FROM " +
		"tracks\n\t\t\tleft join tracks_to_albums"
	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)
	track, err := tracRep.GetTracksByAlbumID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetTrackByGenreID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}
	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes FROM " +
		"tracks\n\t\t\tLEFT JOIN tracks_to_genres"
	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)
	track, err := tracRep.GetTracksByGenreID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetFavoriteTracks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}
	pagination := &commonModels.Pagination{
		Limit:  5,
		Offset: 0,
	}

	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes " +
		"from tracks\n\t\t\tLEFT JOIN tracks_to_user ttu on tracks.track_id = ttu.track_id"
	mock.ExpectQuery(query).WithArgs(uint64(1), pagination.Limit, pagination.Offset).WillReturnRows(rows)
	track, err := tracRep.GetFavoriteTracks(1, pagination)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestCreateTrack(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id",
	}).AddRow(tracksForTests[0].TrackID)
	query := "INSERT"
	mock.ExpectQuery(query).WithArgs(tracksForTests[0].Tittle,
		tracksForTests[0].Text, tracksForTests[0].ReleaseDate).WillReturnRows(rows)
	trackID, err := tracRep.CreateTrack(tracksForTests[0])

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests[0], trackID) {
		t.Fatalf("Not equal")
	}
}

func TestAddTrackToMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "INSERT INTO tracks_to_user"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.AddTrackToMediateka(1, 2)

	assert.NoError(t, err)
}

func TestDeleteTrackFromMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "DELETE FROM tracks_to_user"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.DeleteTrackFromMediateka(1, 2)

	assert.NoError(t, err)
}

func TestAddTrackToFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "UPDATE tracks_to_user SET favorite = true"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.AddTrackToFavorites(1, 2)

	assert.NoError(t, err)
}

func TestDeleteTrackFromFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "UPDATE tracks_to_user SET favorite = false"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.DeleteTrackFromFavorites(1, 2)

	assert.NoError(t, err)
}

func TestUploadAudio(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "UPDATE tracks SET audio"
	mock.ExpectExec(query).WithArgs("/api/v1/data/audio/1.mp3", 1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.UploadAudio(1, "/api/v1/data/audio/1.mp3")

	assert.NoError(t, err)
}

func TestUploadPicture(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "UPDATE tracks SET picture"
	mock.ExpectExec(query).WithArgs("/api/v1/data/img/tracks/3.png", 1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.UploadPicture(1, "/api/v1/data/img/tracks/3.png")

	assert.NoError(t, err)
}

func TestSearchTracks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}

	query := "SELECT track_id, tittle, text, audio, picture, release_date, duration, likes"
	mock.ExpectQuery(query).WithArgs("song").WillReturnRows(rows)
	track, err := tracRep.SearchTracks("song")

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetTop20Tracks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}

	query := "select track_id, tittle, text, audio, picture, release_date, " +
		"duration, likes from tracks\n\t\t\torder by rating desc\n\t\t\tlimit 20"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.GetTop20Tracks()

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetBillbordTopCharts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}
	query := "select track_id, tittle, text, audio, picture, release_date, duration, " +
		"likes from tracks\n\t\t\torder by amount_of_listens desc\n\t\t\tlimit 20"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.GetBillbordTopCharts()

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestGetHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}

	query := "select t.track_id, t.tittle, t.text, t.audio, t.picture, t.release_date, " +
		"t.duration, t.likes\n\t\t\tfrom history as h\n\t\t\tleft join tracks as t on t.track_id " +
		"= h.track_id\n\t\t\twhere h.user_id ="
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	track, err := tracRep.GetHistory(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestAddToHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()
	query := "insert into history"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.AddToHistory(1, 2)

	assert.NoError(t, err)
}

func TestGetMusicianByTrackID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"musician_id", "name", "description", "picture",
	})
	for _, row := range musiciansForTests {
		rows.AddRow(row.MusicianID, row.Name, row.Description,
			row.Picture)
	}
	query := "select mus.musician_id, mus.name, mus.description, mus.picture " +
		"from musicians as mus\n\t\t\t\t\t\tleft join musicians_to_tracks mtt on " +
		"mus.musician_id = mtt.musician_id\n\t\t\t\t\t\twhere mtt.track_id ="
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	musicians := tracRep.GetMusicianByTrackID(1)

	if !reflect.DeepEqual(musiciansForTests, musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetAlbumsByTrackID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})
	for _, row := range albumsForTests {
		rows.AddRow(row.AlbumID, row.Tittle, row.Picture, row.ReleaseDate)
	}
	query := "select a.album_id, a.tittle, a.picture, a.release_date from albums " +
		"as a\n\t\t\t\t\t\tleft join tracks_to_albums tta on a.album_id = " +
		"tta.album_id\n\t\t\t\t\t\twhere tta.track_id = "
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	albums := tracRep.GetAlbumsByTrackID(1)

	if !reflect.DeepEqual(albumsForTests, albums) {
		t.Fatalf("Not equal")
	}
}

func TestGetGenreByTrackID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"genre_id", "title",
	})
	for _, row := range genresForTests {
		rows.AddRow(row.GenreID, row.Title)
	}
	query := "select g.genre_id, g.title from genres as g\n\t\t\t\t\tleft " +
		"join tracks_to_genres ttg on g.genre_id = " +
		"ttg.genre_id\n\t\t\t\t\twhere ttg.track_id = "
	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
	genres := tracRep.GetGenreByTrackID(1)

	if !reflect.DeepEqual(genresForTests, genres) {
		t.Fatalf("Not equal")
	}
}

func TestGetTopTrack(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date", "duration", "likes",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate, row.Duration, row.Likes)
	}

	query := "select track_id, tittle, text, audio, picture, release_date, duration, likes " +
		"from tracks\n\t\t\torder by rating desc\n\t\t\tlimit 1"
	mock.ExpectQuery(query).WillReturnRows(rows)
	track, err := tracRep.GetTopTrack()

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}

func TestIncrementLikes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "update tracks set likes"
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.IncrementLikes(1)

	assert.NoError(t, err)
}

func TestDecrementLikes(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "update tracks set likes = likes - 1"
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.DecrementLikes(1)

	assert.NoError(t, err)
}

func TestCheckTrackInMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "update tracks set likes"

	mock.ExpectExec(query).WithArgs(1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.CheckTrackInMediateka(1, 1)

	assert.Equal(t, err, errors.New("no track"))
}

func TestCheckTrackInFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()
	query := "update tracks set likes"

	mock.ExpectExec(query).WithArgs(1).WillReturnResult(
		sqlmock.NewResult(1, 1))
	err = tracRep.CheckTrackInFavorite(1, 1)

	assert.Equal(t, err, errors.New("no track"))
}
