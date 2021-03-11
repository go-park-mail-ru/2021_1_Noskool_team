package repository

import (
	"2021_1_Noskool_team/internal/app/tracks/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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
