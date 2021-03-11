package repository

import (
	"2021_1_Noskool_team/internal/app/album/models"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
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
