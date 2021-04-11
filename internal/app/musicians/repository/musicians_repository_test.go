package repository

import (
	"2021_1_Noskool_team/internal/app/musicians/models"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedMusician = []models.Musician{
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

	expectedMusicianPointers = []*models.Musician{
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
)

func TestGetMusiciansByGenre(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
	for _, musician := range expectedMusician {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs("rok").WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusiciansByGenre("rok")

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
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

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
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

func TestSearchMusicians(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
	for _, musician := range expectedMusicianPointers[:1] {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "SELECT musician_id, name, description, picture FROM " +
		"musicians WHERE musicians.name LIKE"

	mock.ExpectQuery(query).WithArgs("J").WillReturnRows(rows)

	musicians, err := musiciansRep.SearchMusicians("J")

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusicianPointers[:1], musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetMusicianByTrackID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
	for _, musician := range expectedMusician {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusicianByTrackID(1)

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetMusicianByAlbumID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
	for _, musician := range expectedMusician {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs(2).WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusicianByAlbumID(2)

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetMusicianByPlaylistID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"musician_id", "name", "description", "picture"})
	for _, musician := range expectedMusician {
		rows.AddRow(musician.MusicianID, musician.Name,
			musician.Description, musician.Picture)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusicianByPlaylistID(1)

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
		t.Fatalf("Not equal")
	}
}
