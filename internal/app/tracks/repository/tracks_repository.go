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

func (trackRep *TracksRepository) GetTrackByID(trackID int) (*models.Track, error) {
	track := &models.Track{}
	err := trackRep.con.QueryRow(
		"SELECT * FROM tracks where track_id = $1", trackID,
	).Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
		&track.ReleaseDate)

	return track, err
}

func (trackRep *TracksRepository) CreateTrack(track *models.Track) (*models.Track, error) {
	query := `INSERT INTO tracks (tittle, text, release_date) VALUES
			($1, $2, $3) returning track_id`

	err := trackRep.con.QueryRow(query, track.Tittle,
		track.Text, track.ReleaseDate).Scan(&track.TrackID)
	if err != nil {
		return nil, err
	}
	return track, nil
}

func (trackRep *TracksRepository) UploadAudio(trackID int, audioPath string) error {
	query := `UPDATE tracks SET audio = $1
			WHERE track_id = $2`

	res, err := trackRep.con.Exec(query, audioPath, trackID)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (trackRep *TracksRepository) UploadPicture(trackID int, audioPath string) error {
	query := `UPDATE tracks SET picture = $1
			WHERE track_id = $2`

	res, err := trackRep.con.Exec(query, audioPath, trackID)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}

func (trackRep *TracksRepository) GetTracksByTittle(trackTittle string) ([]*models.Track, error) {
	rows, err := trackRep.con.Query(
		"SELECT * FROM tracks where tittle = $1", trackTittle)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tracksByTittle := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		_ = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)

		tracksByTittle = append(tracksByTittle, track)
	}

	return tracksByTittle, err
}

func (trackRep *TracksRepository) GetTrackByMusicianID(musicianID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tracks.tittle, tracks.text,
       		tracks.audio,tracks.picture, tracks.release_date
			FROM tracks
			INNER JOIN musicians_to_tracks
			ON tracks.track_id = musicians_to_tracks.track_id
			inner join musicians as m on musicians_to_tracks.musician_id = m.musician_id
			where m.musician_id = $1`

	rows, err := trackRep.con.Query(
		query, musicianID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tracksByMusName := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		_ = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)

		tracksByMusName = append(tracksByMusName, track)
	}

	return tracksByMusName, err
}
