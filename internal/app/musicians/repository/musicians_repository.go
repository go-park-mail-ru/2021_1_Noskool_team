package repository

import (
	"2021_1_Noskool_team/internal/app/musicians"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"database/sql"
	_ "github.com/lib/pq"
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

func (musicRep *MusicianRepository) GetMusiciansByGenres(genre string) ([]models.Musician, error) {
	query := `select musicians.* from musicians
		left join musicians_to_genres as m_g on m_g.musician_id = musicians.musician_id
		left join genres on genres.genre_id = m_g.genre_id
		where genres.title = $1`
	musiciansRows, err := musicRep.con.Query(query, genre)
	if err != nil {
		logrus.Error(err)
	}
	defer musiciansRows.Close()
	musicians := make([]models.Musician, 0)

	for musiciansRows.Next() {
		musician := models.Musician{}
		err = musiciansRows.Scan(&musician.MusicianID, &musician.Name,
			&musician.Description, &musician.Picture)

		if err != nil {
			logrus.Error(err)
		}
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
