package repository

import (
	"2021_1_Noskool_team/internal/app/album"
	"2021_1_Noskool_team/internal/app/album/models"
	"database/sql"
	"errors"
	"fmt"

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

func (albumRep *AlbumsRepository) GetAlbumByID(albumID int) (*models.Album, error) {
	album := &models.Album{}
	err := albumRep.con.QueryRow(
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
	query := `select albums.album_id, albums.tittle, albums.picture, albums.release_date from albums 
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

func (albumRep *AlbumsRepository) AddAlbumToFavorites(userID, albumID int) error {
	query := `UPDATE album_to_user SET favorite = true
			WHERE user_id = $1 and album_id = $2`

	res, err := albumRep.con.Exec(query, userID, albumID)
	logrus.Info(res)
	return err
}

func (albumRep *AlbumsRepository) DelteAlbumFromFavorites(userID, albumID int) error {
	query := `UPDATE album_to_user SET favorite = false
			WHERE user_id = $1 and album_id = $2`

	res, err := albumRep.con.Exec(query, userID, albumID)
	logrus.Info(res)
	return err
}

func (albumRep *AlbumsRepository) AddAlbumToMediateka(userID, albumID int) error {
	query := `INSERT INTO album_to_user(user_id, album_id) VALUES ($1, $2);`
	res, err := albumRep.con.Exec(query, userID, albumID)
	fmt.Println(res)
	return err
}

func (albumRep *AlbumsRepository) DeleteAlbumFromMediateka(userID, albumID int) error {
	query := `DELETE FROM album_to_user
			WHERE user_id = $1 and album_id = $2`
	res, err := albumRep.con.Exec(query, userID, albumID)
	fmt.Println(res)
	return err
}

func (albumRep *AlbumsRepository) GetFavoriteAlbums(userID int) ([]*models.Album, error) {
	query := `SELECT a.album_id, a.tittle, a.picture, a.release_date from albums as a
			left join album_to_user atu on a.album_id = atu.album_id
			where atu.user_id = $1 and atu.favorite = true
			order by atu.album_id`
	rows, err := albumRep.con.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	albums := make([]*models.Album, 0)

	for rows.Next() {
		album := &models.Album{}
		err = rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		albums = append(albums, album)
	}
	return albums, err
}

func (albumRep *AlbumsRepository) GetAlbumsMediateka(userID int) ([]*models.Album, error) {
	query := `SELECT a.album_id, a.tittle, a.picture, a.release_date from albums as a
			left join album_to_user atu on a.album_id = atu.album_id
			where atu.user_id = $1
			order by atu.album_id`
	rows, err := albumRep.con.Query(query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	albums := make([]*models.Album, 0)

	for rows.Next() {
		album := &models.Album{}
		err = rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		albums = append(albums, album)
	}
	return albums, err
}

func (albumRep *AlbumsRepository) CheckAlbumInMediateka(userID, albumID int) error {
	query := `select count(*) from album_to_user
	where album_id = $1 and user_id = $2`

	res := 0
	err := albumRep.con.QueryRow(query, userID, albumID).Scan(&res)
	if res < 1 {
		return errors.New("no track")
	}
	return err
}

func (albumRep *AlbumsRepository) CheckAlbumInFavorite(userID, albumID int) error {
	query := `select count(*)from album_to_user
	where album_id = $1 and user_id = $2 and favorite = true`

	res := 0
	err := albumRep.con.QueryRow(query, userID, albumID).Scan(&res)
	if res < 1 {
		return errors.New("no track")
	}
	return err
}

func (albumRep *AlbumsRepository) GetAlbums() ([]*models.Album, error) {
	query := `select album_id, tittle, picture, release_date
			  from albums
			  order by rating desc
			  limit 20`
	rows, err := albumRep.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := make([]*models.Album, 0)

	for rows.Next() {
		tmp := &models.Album{}
		err = rows.Scan(
			&tmp.AlbumID,
			&tmp.Tittle,
			&tmp.Picture,
			&tmp.ReleaseDate)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
