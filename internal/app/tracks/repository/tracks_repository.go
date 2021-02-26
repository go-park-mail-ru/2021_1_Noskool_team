package repository

import (
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	"database/sql"
	"fmt"
)

type TracksRepository struct {
	con *sql.DB
}

func NewTracksRepository(con *sql.DB) tracks.Repository {
	return &TracksRepository{
		con: con,
	}
}

func (trackRep *TracksRepository) GetTrackById(trackId int) (*models.Track, error) {
	track := &models.Track{}
	err := trackRep.con.QueryRow(
		"SELECT * FROM tracks where track_id = $1", trackId,
	).Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio,&track.Picture,
		&track.ReleaseDate)
	return track, err
}

func (trackRep *TracksRepository) GetTracksByTittle(trackTittle string) ([]*models.Track, error) {
	rows, err := trackRep.con.Query(
		"SELECT * FROM tracks where tittle = $1", trackTittle)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		_ = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio,&track.Picture,
			&track.ReleaseDate)
		tracks = append(tracks, track)
	}
	return tracks, err
}

