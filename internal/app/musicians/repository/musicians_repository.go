package repository

import (
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type MusicianRepository struct {
	con *sql.DB
}

func NewMusicRepository(con *sql.DB) musicians.Repository {
	return &MusicianRepository{
		con: con,
	}
}

func (musicRep *MusicianRepository) GetMusiciansByGenres(genre string) ([]models.Musician, error) {
	var genreID int
	err := musicRep.con.QueryRow(
		"SELECT genre_id FROM genres where title = $1", genre,
	).Scan(&genreID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := musicRep.con.Query(
		"SELECT musician_id FROM musicians_to_genres where genre_id = $1",
		genreID)

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

		fmt.Println(err)
		musicians = append(musicians, musician)
	}
	return musicians, nil
}

func (musicRep *MusicianRepository) GetMusicianByID(musicianID int) (*models.Musician, error) {
	musician := &models.Musician{}
	err := musicRep.con.QueryRow(
		"SELECT * FROM musicians where musician_id = $1", musicianID,
	).Scan(&musician.MusicianID, &musician.Name, &musician.Description, &musician.Picture)

	return musician, err
}
