package repository

import (
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
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

func (trackRep *TracksRepository) GetTracksByUserID(userID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date from tracks
			LEFT JOIN tracks_to_user ttu on tracks.track_id = ttu.track_id
			where ttu.user_id = $1`

	rows, err := trackRep.con.Query(
		query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		tracks = append(tracks, track)
	}
	return tracks, err
}

func (trackRep *TracksRepository) GetFavoriteTracks(userID int,
	pagination *commonModels.Pagination) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date from tracks
			LEFT JOIN tracks_to_user ttu on tracks.track_id = ttu.track_id
			where ttu.user_id = $1 and ttu.favorite = true
			order by tracks.track_id
			limit $2
			offset $3`

	rows, err := trackRep.con.Query(
		query, userID, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		tracks = append(tracks, track)
	}
	return tracks, err
}

func (trackRep *TracksRepository) AddTrackToFavorites(userID, trackID int) error {
	query := `UPDATE tracks_to_user SET favorite = true
			WHERE user_id = $1 and track_id = $2`

	res, err := trackRep.con.Exec(query, userID, trackID)
	logrus.Info(res)
	return err
}

func (trackRep *TracksRepository) DeleteTrackFromFavorites(userID, trackID int) error {
	query := `UPDATE tracks_to_user SET favorite = false
			WHERE user_id = $1 and track_id = $2`

	res, err := trackRep.con.Exec(query, userID, trackID)
	logrus.Info(res)
	return err
}

func (trackRep *TracksRepository) GetTracksByAlbumID(albumID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date FROM tracks
			left join tracks_to_albums tta on tracks.track_id = tta.track_id
			WHERE tta.album_id = $1`
	rows, err := trackRep.con.Query(query, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracksByAlbum := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		tracksByAlbum = append(tracksByAlbum, track)
	}
	return tracksByAlbum, nil
}

func (trackRep *TracksRepository) GetTracksByGenreID(genreID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date FROM tracks
			LEFT JOIN tracks_to_genres ttg ON tracks.track_id = ttg.track_id
			WHERE ttg.genre_id = $1`
	rows, err := trackRep.con.Query(query, genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracksByGenre := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			logrus.Error(err)
		}
		tracksByGenre = append(tracksByGenre, track)
	}
	return tracksByGenre, nil
}

func (trackRep *TracksRepository) AddTrackToMediateka(userID, trackID int) error {
	query := `INSERT INTO tracks_to_user(user_id, track_id) VALUES ($1, $2);`
	res, err := trackRep.con.Exec(query, userID, trackID)
	fmt.Println(res)
	return err
}

func (trackRep *TracksRepository) DeleteTrackFromMediateka(userID, trackID int) error {
	query := `DELETE FROM tracks_to_user
			WHERE user_id = $1 and track_id = $2`
	res, err := trackRep.con.Exec(query, userID, trackID)
	fmt.Println(res)
	return err
}

func (trackRep *TracksRepository) SearchTracks(searchQuery string) ([]*models.Track, error) {
	query := `SELECT track_id, tittle, text, audio, picture, release_date FROM tracks
			WHERE tittle LIKE '%' || $1 || '%'`

	rows, err := trackRep.con.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracksByQuery := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			return nil, err
		}
		tracksByQuery = append(tracksByQuery, track)
	}
	return tracksByQuery, nil
}
