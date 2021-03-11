package repository

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetMusiciansByGenres(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	expectedMusician := &[]models.Musician{
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
	}

	rows := sqlmock.NewRows([]string{
		"musician_id", "name", "description", "picture",
	})
	for _, musician := range *expectedMusician {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs("rok").WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusiciansByGenres("rok")

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetMusicianByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	expectedMusician := &models.Musician{
		MusicianID:  1,
		Name:        "Joji",
		Description: "Pretty Boy",
		Picture:     "picture",
	}

	rows := sqlmock.NewRows([]string{
		"musician_id", "name", "description", "picture",
	})
	rows.AddRow(expectedMusician.MusicianID, expectedMusician.Name,
		expectedMusician.Description, expectedMusician.Picture)
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusicianByID(1)

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, musicians) {
		t.Fatalf("Not equal")
	}
}
