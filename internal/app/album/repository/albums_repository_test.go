package repository

import (
	"2021_1_Noskool_team/internal/app/album/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedAlbums = &[]models.Album{
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
	}
	expectedAlbumspointers = []*models.Album{
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
)

func TestGetAlbumByID(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	albumsRep := NewAlbumsRepository(dbCon)
	defer dbCon.Close()

	expectedAlbum := &models.Album{
		AlbumID:     1,
		Tittle:      "album",
		Picture:     "picture",
		ReleaseDate: "date",
	}

	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})
	rows.AddRow(expectedAlbum.AlbumID, expectedAlbum.Tittle,
		expectedAlbum.Picture, expectedAlbum.ReleaseDate)
	query := "select"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	album, err := albumsRep.GetAlbumByID(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbum, album) {
		t.Fatalf("Not equal")
	}
}

func TestSearchAlbums(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	albumsRep := NewAlbumsRepository(dbCon)

	defer dbCon.Close()
	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})

	for _, row := range expectedAlbumspointers {
		rows.AddRow(row.AlbumID, row.Tittle,
			row.Picture, row.ReleaseDate)
	}
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs("album").WillReturnRows(rows)

	album, err := albumsRep.SearchAlbums("album")
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbumspointers, album) {
		t.Fatalf("Not equal")
	}
}

func TestGetAlbumsByMusicianID(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	albumsRep := NewAlbumsRepository(dbCon)

	defer dbCon.Close()
	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})

	for _, row := range *expectedAlbums {
		rows.AddRow(row.AlbumID, row.Tittle,
			row.Picture, row.ReleaseDate)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	album, err := albumsRep.GetAlbumsByMusicianID(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbums, album) {
		t.Fatalf("Not equal")
	}
}

func TestGetAlbumsByTrackID(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	albumsRep := NewAlbumsRepository(dbCon)

	defer dbCon.Close()
	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})

	for _, row := range *expectedAlbums {
		rows.AddRow(row.AlbumID, row.Tittle,
			row.Picture, row.ReleaseDate)
	}
	query := "select"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	album, err := albumsRep.GetAlbumsByTrackID(1)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbums, album) {
		t.Fatalf("Not equal")
	}
}

func TestAddAlbumToFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewAlbumsRepository(db)
	defer db.Close()
	query := "UPDATE album_to_user SET favorite = true"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))

	err = tracRep.AddAlbumToFavorites(1, 2)

	assert.NoError(t, err)
}

func TestDelteAlbumFromFavorites(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewAlbumsRepository(db)
	defer db.Close()
	query := "UPDATE album_to_user SET favorite = false"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))

	err = tracRep.DelteAlbumFromFavorites(1, 2)

	assert.NoError(t, err)
}

func TestAddAlbumToMediateka(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewAlbumsRepository(db)
	defer db.Close()
	query := "INSERT INTO album_to_user"
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))

	err = tracRep.AddAlbumToMediateka(1, 2)

	assert.NoError(t, err)
}

func TestDeleteAlbumFromMediatekaa(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	tracRep := NewAlbumsRepository(db)
	defer db.Close()
	query := "DELETE FROM album_to_user\n\t\t\tWHERE user_id ="
	mock.ExpectExec(query).WithArgs(1, 2).WillReturnResult(
		sqlmock.NewResult(1, 1))

	err = tracRep.DeleteAlbumFromMediateka(1, 2)

	assert.NoError(t, err)
}

func TestGetFavoriteAlbums(t *testing.T) {
	dbCon, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock '%s'", err)
	}
	albumsRep := NewAlbumsRepository(dbCon)

	defer dbCon.Close()
	rows := sqlmock.NewRows([]string{
		"album_id", "tittle", "picture", "release_date",
	})
	for _, row := range expectedAlbumspointers {
		rows.AddRow(row.AlbumID, row.Tittle, row.Picture, row.ReleaseDate)
	}
	query := "SELECT a.album_id, a.tittle, a.picture, a.release_date from " +
		"albums as a\n\t\t\tleft join album_to_user atu on a.album_id = " +
		"atu.album_id\n\t\t\twhere atu.user_id ="
	pagination := &commonModels.Pagination{
		Limit:  1,
		Offset: 2,
	}

	mock.ExpectQuery(query).WithArgs(1, pagination.Limit,
		pagination.Offset).WillReturnRows(rows)

	album, err := albumsRep.GetFavoriteAlbums(1, pagination)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbumspointers, album) {
		t.Fatalf("Not equal")
	}
}
