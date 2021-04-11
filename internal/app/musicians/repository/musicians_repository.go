package repository

import (
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"database/sql"

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
	query := `select musicians.* from musicians
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
		"SELECT * FROM musicians where musician_id = $1", musicianID).Scan(
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
	query := `select musicians.* from musicians
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
	query := `select musicians.* from musicians
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
	query := `select musicians.* from musicians
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
