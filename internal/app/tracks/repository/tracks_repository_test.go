package repository

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
