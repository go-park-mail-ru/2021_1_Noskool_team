package repository

import (
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/album/models"
	"database/sql"

	"github.com/sirupsen/logrus"
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
		"select album_id, tittle, picture, release_date FROM albums where album_id = $1", albumID).Scan(
		&album.AlbumID,
		&album.Tittle,
		&album.Picture,
		&album.ReleaseDate)
	if err != nil {
		logrus.Error("error in db GetAlbumByID: ", err)
		return nil, err
	}
	return album, nil
}

func (albumRep *AlbumsRepository) GetAlbumsByMusicianID(musicianID int) (*[]models.Album, error) {
	query := `select album_id, tittle, picture, release_date from albums 
	left join Musicians_to_Albums as m_a on m_a.album_id = albums.album_id 
	left join musicians on musicians.musician_id = m_a.musician_id 
	where musicians.musician_id = $1`
	rows, err := albumRep.con.Query(query, musicianID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albumsByQuery := make([]models.Album, 0)
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			return nil, err
		}
		albumsByQuery = append(albumsByQuery, album)
	}
	return &albumsByQuery, nil
}

func (albumRep *AlbumsRepository) GetAlbumsByTrackID(trackID int) (*[]models.Album, error) {
	query := `select album_id, tittle, picture, release_date from albums 
	left join Tracks_to_Albums as t_a on t_a.album_id = albums.album_id 
	left join tracks on tracks.track_id = t_a.track_id  
	where tracks.track_id = $1`
	rows, err := albumRep.con.Query(query, trackID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albumsByQuery := make([]models.Album, 0)
	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			return nil, err
		}
		albumsByQuery = append(albumsByQuery, album)
	}
	return &albumsByQuery, nil
}

func (albumRep *AlbumsRepository) SearchAlbums(searchQuery string) (*[]models.Album, error) {
	query := `SELECT album_id, tittle, picture, release_date FROM albums 
			WHERE tittle LIKE '%' || $1 || '%'`
	rows, err := albumRep.con.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	albumsByQuery := make([]models.Album, 0)

	for rows.Next() {
		album := models.Album{}
		err := rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			return nil, err
		}
		albumsByQuery = append(albumsByQuery, album)
	}
	return &albumsByQuery, nil
}
