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

func (albumRep *AlbumsRepository) SearchAlbums(searchQuery string) ([]*models.Album, error) {
	query := `SELECT album_id, tittle, picture, release_date FROM albums
			WHERE tittle LIKE '%' || $1 || '%'`

	rows, err := albumRep.con.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albumsByQuery := make([]*models.Album, 0)

	for rows.Next() {
		album := &models.Album{}
		err := rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			return nil, err
		}
		albumsByQuery = append(albumsByQuery, album)
	}
	return albumsByQuery, nil
}
