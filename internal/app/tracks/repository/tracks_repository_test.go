package repository

import (
	"2021_1_Noskool_team/internal/app/tracks/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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
			Tittle:      "helloWorld",
			Text:        "sing song song ooooo",
			Audio:       "/api/v1/data/audio/2.mp3",
			Picture:     "/api/v1/data/audio/tracks/2.mp3",
			ReleaseDate: "2020-03-04",
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
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	}).AddRow(1, "song", "sing song song", "audio", "picture", "date")
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
	}

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	}).AddRow(expectedTrack.TrackID, expectedTrack.Tittle, expectedTrack.Text,
		expectedTrack.Audio, expectedTrack.Picture, expectedTrack.ReleaseDate)
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
	}

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	}).AddRow(expectedTrack.TrackID, expectedTrack.Tittle, expectedTrack.Text,
		expectedTrack.Audio, expectedTrack.Picture, expectedTrack.ReleaseDate)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs("hello").WillReturnRows(rows)

	track, err := tracRep.GetTracksByTittle("hello")

	fmt.Println(track)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedTrack, track[0]) {
		t.Fatalf("Not equal")
	}
}

func TestGetTrackByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewTracksRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate)
	}
	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date from " +
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
		"track_id", "tittle", "text", "audio", "picture", "release_date",
	})
	for _, row := range tracksForTests {
		rows.AddRow(row.TrackID, row.Tittle, row.Text,
			row.Audio, row.Picture, row.ReleaseDate)
	}
	query := "SELECT tracks.track_id, tittle, text, audio, picture, release_date FROM " +
		"tracks\n\t\t\tleft join tracks_to_albums"
	mock.ExpectQuery(query).WithArgs(uint64(1)).WillReturnRows(rows)
	track, err := tracRep.GetTracksByAlbumID(1)

	assert.NoError(t, err)
	if !reflect.DeepEqual(tracksForTests, track) {
		t.Fatalf("Not equal")
	}
}
