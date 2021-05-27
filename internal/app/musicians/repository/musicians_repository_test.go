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

func TestGetMusiciansTop4(t *testing.T) {
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

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusiciansTop4()

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetMusicians(t *testing.T) {
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

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	musicians, err := musiciansRep.GetMusicians()

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedMusician, *musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetGenreForMusician(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"title"})
	rows.AddRow("rok")

	query := "select"

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	musicians, err := musiciansRep.GetGenreForMusician("Nirvana")

	fmt.Println(musicians)
	assert.NoError(t, err)
	if !reflect.DeepEqual([]string{"rok"}, *musicians) {
		t.Fatalf("Not equal")
	}
}

func TestGetGenreForMusicianFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)

	defer db.Close()

	rows := sqlmock.NewRows([]string{"title"})
	rows.AddRow("rok")

	query := "select"

	mock.ExpectQuery(query).WithArgs().WillReturnError(fmt.Errorf("Ошибка получения жанров музыканта: %s", "Nirvana"))

	musicians, err := musiciansRep.GetGenreForMusician("Nirvana")

	fmt.Println(musicians)
	assert.Equal(t, err, fmt.Errorf("Ошибка получения жанров музыканта: %s", "Nirvana"))
}

func TestAddMusicianToMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	query := "INSERT"
	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = musiciansRep.AddMusicianToMediateka(1, 1)
	assert.NoError(t, err)
}

func TestDeleteMusicianFromMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	query := "DELETE"
	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = musiciansRep.DeleteMusicianFromMediateka(1, 1)
	assert.NoError(t, err)
}

func TestCheckMusicianInMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"user_id", "musician_id"})
	rows.AddRow(1, 1)

	query := "select"

	mock.ExpectQuery(query).WithArgs(1, 1).WillReturnRows(rows)
	err = musiciansRep.CheckMusicianInMediateka(1, 1)
	assert.Equal(t, err, fmt.Errorf("no musician"))
}

func TestAddMusicianToFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	query := "UPDATE"
	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = musiciansRep.AddMusicianToFavorites(1, 1)
	assert.NoError(t, err)
}

func TestDeleteMusicianFromFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	query := "UPDATE"
	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = musiciansRep.DeleteMusicianFromFavorites(1, 1)
	assert.NoError(t, err)
}

func TestGetMusiciansMediateka(t *testing.T) {
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

	musicians, err := musiciansRep.GetMusiciansMediateka(2)

	fmt.Println(musicians)
	assert.NoError(t, err)
}

func TestGetMusiciansFavorites(t *testing.T) {
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

	musicians, err := musiciansRep.GetMusiciansFavorites(2)

	fmt.Println(musicians)
	assert.NoError(t, err)
}

func TestCheckMusicianInFavorite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	musiciansRep := NewMusicRepository(db)
	defer db.Close()

	query := "UPDATE"
	mock.ExpectExec(query).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	err = musiciansRep.CheckMusicianInFavorite(1, 1)
	assert.Equal(t, err, fmt.Errorf("no musician"))
}
