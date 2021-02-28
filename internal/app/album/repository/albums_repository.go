package repository

import (
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/album/models"
	"database/sql"
)

type AlbumsRepository struct {
	con *sql.DB
}

func NewAlbumsRepository(con *sql.DB) album.Repository {
	return &AlbumsRepository{
		con: con,
	}
}

func (albumsRep *AlbumsRepository) GetAlbumByID(albumID int) (*models.Album, error) {
	album := &models.Album{}
	err := albumsRep.con.QueryRow(
		"SELECT * FROM albums where album_id = $1", albumID,
	).Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)

	return album, err
}
