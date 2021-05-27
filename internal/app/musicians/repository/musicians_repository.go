package repository

import (
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" //goland:noinspection GoLinterLocal
	"github.com/sirupsen/logrus"
)

type MusicianRepository struct {
	con *sql.DB
}

func NewMusicRepository(con *sql.DB) musicians.Repository {
	return &MusicianRepository{
		con: con,
	}
}

func (musicRep *MusicianRepository) GetMusiciansByGenre(genre string) (*[]models.Musician, error) {
	query := `select musicians.musician_id, musicians.name, musicians.description, musicians.picture from musicians 
		left join musicians_to_genres as m_g on m_g.musician_id = musicians.musician_id
		left join genres on genres.genre_id = m_g.genre_id 
		where genres.title = $1`
	musiciansRows, err := musicRep.con.Query(query, genre)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()

	musicians := make([]models.Musician, 0)
	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) GetMusicianByID(musicianID int) (*models.Musician, error) {
	musician := &models.Musician{}
	err := musicRep.con.QueryRow(
		"SELECT musician_id, name, description, picture FROM musicians where musician_id = $1", musicianID).Scan(
		&musician.MusicianID,
		&musician.Name,
		&musician.Description,
		&musician.Picture)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return musician, nil
}

func (musicRep *MusicianRepository) GetMusicianByTrackID(trackID int) (*[]models.Musician, error) {
	query := `select musicians.musician_id, musicians.name, musicians.description, musicians.picture from musicians 
		left join Musicians_to_Tracks as m_t on m_t.musician_id = musicians.musician_id
		left join Tracks on Tracks.track_id = m_t.track_id 
		where Tracks.track_id = $1`
	musiciansRows, err := musicRep.con.Query(query, trackID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) GetMusicianByAlbumID(albumID int) (*[]models.Musician, error) {
	query := `select musicians.musician_id, musicians.name, musicians.description, musicians.picture from musicians 
		left join Musicians_to_Albums as m_a on m_a.musician_id = musicians.musician_id
		left join Albums on Albums.album_id = m_a.album_id 
		where Albums.album_id = $1`
	musiciansRows, err := musicRep.con.Query(query, albumID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) GetMusicianByPlaylistID(playlistID int) (*[]models.Musician, error) {
	query := `select musicians.musician_id, musicians.name, musicians.description, musicians.picture from musicians 
		left join Musicians_to_Playlist as m_p on m_p.musician_id = musicians.musician_id
		left join playlists on playlists.playlist_id = m_p.playlist_id  
		where playlists.playlist_id = $1`
	musiciansRows, err := musicRep.con.Query(query, playlistID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) GetMusiciansTop4() (*[]models.Musician, error) {
	musiciansRows, err := musicRep.con.Query("select musician_id, name, description, picture from musicians order by rating limit 4")
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) SearchMusicians(searchQuery string) ([]*models.Musician, error) {
	query := `SELECT musician_id, name, description, picture FROM musicians
			WHERE musicians.name LIKE '%' || $1 || '%'`

	rows, err := musicRep.con.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	musiciansByQuery := make([]*models.Musician, 0)

	for rows.Next() {
		musician := &models.Musician{}
		err := rows.Scan(&musician.MusicianID, &musician.Name, &musician.Description, &musician.Picture)
		if err != nil {
			return nil, err
		}
		musiciansByQuery = append(musiciansByQuery, musician)
	}
	return musiciansByQuery, nil
}

func (musicRep *MusicianRepository) GetMusicians() (*[]models.Musician, error) {
	query := `select musician_id, name, description, picture 
	          from musicians 
			  order by rating desc 
			  limit 20`
	musiciansRows, err := musicRep.con.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return &musicians, nil
}

func (musicRep *MusicianRepository) GetGenreForMusician(nameMusician string) (*[]string, error) {
	query := `select g.title from musicians as m 
	left join Musicians_to_Genres as m_g on (m_g.musician_id = m.musician_id) 
	left join genres as g on (g.genre_id = m_g.genre_id) 
	where m.name = $1;`

	rows, err := musicRep.con.Query(query, nameMusician)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	genres := make([]string, 0)

	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			logrus.Error(err)
			return nil, fmt.Errorf("Ошибка получения жанров музыканта: %s", nameMusician)
		}
		genres = append(genres, tmp)
	}
	return &genres, nil

}

func (musicRep *MusicianRepository) AddMusicianToMediateka(userID, musicianID int) error {
	query := `INSERT INTO musicians_to_user(user_id, musician_id) VALUES ($1, $2)`
	_, err := musicRep.con.Exec(query, userID, musicianID)
	return err
}

func (musicRep *MusicianRepository) DeleteMusicianFromMediateka(userID, musicianID int) error {
	query := `DELETE FROM musicians_to_user
			WHERE user_id = $1 and musician_id = $2`
	res, err := musicRep.con.Exec(query, userID, musicianID)
	fmt.Println(res)
	return err
}

func (musicRep *MusicianRepository) CheckMusicianInMediateka(userID, musicianID int) error {
	query := `select count(*) from musicians_to_user
	where musician_id = $1 and user_id = $2`

	res := 0
	err := musicRep.con.QueryRow(query, musicianID, userID).Scan(&res)
	if res < 1 {
		return errors.New("no musician")
	}
	return err
}

func (musicRep *MusicianRepository) CheckMusicianInFavorite(userID, musicianID int) error {
	query := `select count(*) from musicians_to_user
	where musician_id = $1 and user_id = $2 and favorite = true`

	res := 0
	err := musicRep.con.QueryRow(query, musicianID, userID).Scan(&res)
	if res < 1 {
		return errors.New("no musician")
	}
	return err
}

func (musicRep *MusicianRepository) AddMusicianToFavorites(userID, musicianID int) error {
	query := `UPDATE musicians_to_user SET favorite = true
			WHERE user_id = $1 and musician_id = $2`

	res, err := musicRep.con.Exec(query, userID, musicianID)
	logrus.Info(res)
	return err
}

func (musicRep *MusicianRepository) DeleteMusicianFromFavorites(userID, musicianID int) error {
	query := `UPDATE musicians_to_user SET favorite = false
			WHERE user_id = $1 and musician_id = $2`

	res, err := musicRep.con.Exec(query, userID, musicianID)
	logrus.Info(res)
	return err
}

func (musicRep *MusicianRepository) GetMusiciansMediateka(userID int) ([]*models.Musician, error) {
	query := `select musicians.musician_id, name, description, picture from musicians
			left join musicians_to_user mtu on musicians.musician_id = mtu.musician_id
			where mtu.user_id = $1`
	musiciansRows, err := musicRep.con.Query(query, userID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]*models.Musician, 0)

	for musiciansRows.Next() {
		musician := &models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return musicians, nil
}

func (musicRep *MusicianRepository) GetMusiciansFavorites(userID int) ([]*models.Musician, error) {
	query := `select musicians.musician_id, name, description, picture from musicians
			left join musicians_to_user mtu on musicians.musician_id = mtu.musician_id
			where mtu.user_id = $1 and mtu.favorite = true`
	musiciansRows, err := musicRep.con.Query(query, userID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]*models.Musician, 0)

	for musiciansRows.Next() {
		musician := &models.Musician{}
		err = musiciansRows.Scan(
			&musician.MusicianID,
			&musician.Name,
			&musician.Description,
			&musician.Picture)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		musicians = append(musicians, musician)
	}
	return musicians, nil
}
