package repository

import (
	"2021_1_Noskool_team/internal/app/music"
	"2021_1_Noskool_team/internal/app/music/models"
	"database/sql"
	"fmt"
)

type MusicRepository struct {
	con *sql.DB
}

func NewMusicRepository(con *sql.DB) music.Repository {
	return &MusicRepository{
		con: con,
	}
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
