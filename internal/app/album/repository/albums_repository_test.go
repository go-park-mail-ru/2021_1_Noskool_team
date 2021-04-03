package repository

import (
	"2021_1_Noskool_team/internal/app/album/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	expectedAlbums = []*models.Album{
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
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

	album, err := albumsRep.GetAlbumByID(1)

	fmt.Println(album)
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

	for _, row := range expectedAlbums {
		rows.AddRow(row.AlbumID, row.Tittle,
			row.Picture, row.ReleaseDate)
	}
	query := "SELECT"

	mock.ExpectQuery(query).WithArgs("album").WillReturnRows(rows)

	album, err := albumsRep.SearchAlbums("album")

	fmt.Println(album)
	assert.NoError(t, err)
	if !reflect.DeepEqual(expectedAlbums, album) {
		t.Fatalf("Not equal")
	}
}
