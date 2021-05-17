package repository

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/app/tracks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"database/sql"
	"errors"
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
		"SELECT track_id, tittle, text, audio, picture, release_date, duration, likes FROM tracks where track_id = $1", trackID,
	).Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
		&track.ReleaseDate, &track.Duration, &track.Likes)

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
			&track.ReleaseDate, &track.Duration)

		tracksByTittle = append(tracksByTittle, track)
	}

	return tracksByTittle, err
}

func (trackRep *TracksRepository) GetTrackByMusicianID(musicianID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tracks.tittle, tracks.text,
       		tracks.audio,tracks.picture, tracks.release_date, tracks.duration, tracks.likes
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
			&track.ReleaseDate, &track.Duration, &track.Likes)

		tracksByMusName = append(tracksByMusName, track)
	}

	return tracksByMusName, err
}

func (trackRep *TracksRepository) GetTracksByUserID(userID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes from tracks
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
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			logrus.Error(err)
		}
		tracks = append(tracks, track)
	}
	return tracks, err
}

func (trackRep *TracksRepository) GetFavoriteTracks(userID int,
	pagination *commonModels.Pagination) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes from tracks
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
			&track.ReleaseDate, &track.Duration, &track.Likes)
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
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes FROM tracks
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
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			logrus.Error(err)
		}
		tracksByAlbum = append(tracksByAlbum, track)
	}
	return tracksByAlbum, nil
}

func (trackRep *TracksRepository) GetTracksByGenreID(genreID int) ([]*models.Track, error) {
	query := `SELECT tracks.track_id, tittle, text, audio, picture, release_date, duration, likes FROM tracks
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
			&track.ReleaseDate, &track.Duration, &track.Likes)
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
	query := `SELECT track_id, tittle, text, audio, picture, release_date, duration, likes FROM tracks
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
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			return nil, err
		}
		tracksByQuery = append(tracksByQuery, track)
	}
	return tracksByQuery, nil
}

func (trackRep *TracksRepository) GetTop20Tracks() ([]*models.Track, error) {
	query := `select track_id, tittle, text, audio, picture, release_date, duration, likes from tracks
			order by rating desc
			limit 20`

	rows, err := trackRep.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (trackRep *TracksRepository) GetBillbordTopCharts() ([]*models.Track, error) {
	query := `select track_id, tittle, text, audio, picture, release_date, duration, likes from tracks
			order by amount_of_listens desc
			limit 20`

	rows, err := trackRep.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (trackRep *TracksRepository) GetHistory(userID int) ([]*models.Track, error) {
	query := `select t.track_id, t.tittle, t.text, t.audio, t.picture, t.release_date, t.duration, t.likes
			from history as h
			left join tracks as t on t.track_id = h.track_id
			where h.user_id = $1
			order by creation_date desc`

	rows, err := trackRep.con.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (trackRep *TracksRepository) AddToHistory(userID, trackID int) error {
	query := `insert into history (user_id, track_id) values ($1, $2)`

	_, err := trackRep.con.Exec(query, userID, trackID)
	return err
}

func (trackRep *TracksRepository) CheckTrackInMediateka(userID, trackID int) error {
	query := `select count(*) from tracks_to_user
	where track_id = $1 and user_id = $2`

	res := 0
	err := trackRep.con.QueryRow(query, trackID, userID).Scan(&res)
	if res < 1 {
		return errors.New("no track")
	}
	return err
}

func (trackRep *TracksRepository) GetMusiciansGenresAndAlbums(tracks []*models.Track) []*models.Track {
	for _, track := range tracks {
		track.Musicians = trackRep.GetMusicianByTrackID(track.TrackID)
		track.Albums = trackRep.GetAlbumsByTrackID(track.TrackID)
		track.Genres = trackRep.GetGenreByTrackID(track.TrackID)
	}
	return tracks
}

func (trackRep *TracksRepository) GetMusicianByTrackID(trackID int) []*musiciansModels.Musician {
	queryGetMusicians := `select mus.musician_id, mus.name, mus.description, mus.picture from musicians as mus
						left join musicians_to_tracks mtt on mus.musician_id = mtt.musician_id
						where mtt.track_id = $1`

	rows, err := trackRep.con.Query(queryGetMusicians, trackID)
	if err != nil {
		return nil
	}
	defer rows.Close()
	musicians := make([]*musiciansModels.Musician, 0)

	for rows.Next() {
		musician := &musiciansModels.Musician{}
		err := rows.Scan(&musician.MusicianID, &musician.Name, &musician.Description,
			&musician.Picture)
		if err != nil {
			return nil
		}
		musicians = append(musicians, musician)
	}
	return musicians
}

func (trackRep *TracksRepository) GetAlbumsByTrackID(trackID int) []*albumsModels.Album {
	queryGetAlbums := `select a.album_id, a.tittle, a.picture, a.release_date from albums as a
						left join tracks_to_albums tta on a.album_id = tta.album_id
						where tta.track_id = $1`

	rows, err := trackRep.con.Query(queryGetAlbums, trackID)
	if err != nil {
		return nil
	}
	defer rows.Close()
	albums := make([]*albumsModels.Album, 0)

	for rows.Next() {
		album := &albumsModels.Album{}
		err := rows.Scan(&album.AlbumID, &album.Tittle, &album.Picture, &album.ReleaseDate)
		if err != nil {
			return nil
		}
		albums = append(albums, album)
	}
	return albums
}

func (trackRep *TracksRepository) GetGenreByTrackID(trackID int) []*commonModels.Genre {
	queryGetGenre := `select g.genre_id, g.title from genres as g
					left join tracks_to_genres ttg on g.genre_id = ttg.genre_id
					where ttg.track_id = $1`
	rows, err := trackRep.con.Query(queryGetGenre, trackID)
	if err != nil {
		return nil
	}
	defer rows.Close()
	genres := make([]*commonModels.Genre, 0)

	for rows.Next() {
		genre := &commonModels.Genre{}
		err := rows.Scan(&genre.GenreID, &genre.Title)
		if err != nil {
			return nil
		}
		genres = append(genres, genre)
	}
	return genres
}

func (trackRep *TracksRepository) GetTopTrack() ([]*models.Track, error) {
	query := `select track_id, tittle, text, audio, picture, release_date, duration, likes from tracks
			order by rating desc
			limit 1`

	rows, err := trackRep.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracks := make([]*models.Track, 0)

	for rows.Next() {
		track := &models.Track{}
		err := rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate, &track.Duration, &track.Likes)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (trackRep *TracksRepository) CheckTrackInFavorite(userID, trackID int) error {
	query := `select count(*) from tracks_to_user
	where track_id = $1 and user_id = $2 and favorite = true`

	res := 0
	err := trackRep.con.QueryRow(query, trackID, userID).Scan(&res)
	if res < 1 {
		return errors.New("no track")
	}
	return err
}

func (trackRep *TracksRepository) IncrementLikes(trackID int) error {
	query := `update tracks set likes = likes + 1
			where music_service.public.tracks.track_id = $1`

	_, err := trackRep.con.Exec(query, trackID)
	return err
}

func (trackRep *TracksRepository) DecrementLikes(trackID int) error {
	query := `update tracks set likes = likes - 1
			where music_service.public.tracks.track_id = $1`

	_, err := trackRep.con.Exec(query, trackID)
	return err
}
