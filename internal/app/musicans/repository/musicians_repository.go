package repository

import (
	"2021_1_Noskool_team/internal/app/musicans"
	"2021_1_Noskool_team/internal/app/musicans/models"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type MusicRepository struct {
	con *sql.DB
}

func NewMusicRepository(con *sql.DB) musicans.Repository {
	return &MusicRepository{
		con: con,
	}
}

func (musicRep *MusicRepository) GetMusiciansByGenres(genre string) ([]models.Musician, error) {
	var genreID int
	err := musicRep.con.QueryRow(
		"SELECT genre_id FROM genres where title = $1", genre,
	).Scan(&genreID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := musicRep.con.Query("SELECT musician_id FROM musicians_to_genres where genre_id = $1", genreID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	musiciansIDs := make([]int, 0)

	for rows.Next() {
		var musID int
		_ = rows.Scan(&musID)
		musiciansIDs = append(musiciansIDs, musID)
	}
	preperedQuery, err := musicRep.con.Prepare("SELECT * FROM musicians WHERE musician_id = ANY ($1)")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	musiciansRows, err := preperedQuery.Query(pq.Array(musiciansIDs))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)
	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(&musician.MusicianID, &musician.Name,
			&musician.Description, &musician.Picture)
		musicians = append(musicians, musician)
	}
	return musicians, nil
}

func (musicRep *MusicRepository) GetTrackById(trackId int) (models.Track, error) {
	track := models.Track{}
	err := musicRep.con.QueryRow(
		"SELECT song_id, tittle, text, picture, release_date FROM songs where song_id = $1", trackId,
	).Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Picture,
		&track.ReleaseDate)
	fmt.Println(err)
	return track, err
}
