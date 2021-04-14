package repository

import (
	"2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"database/sql"
	"fmt"
)

type PlaylistRepository struct {
	con *sql.DB
}

func NewPlaylistRepository(newCon *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{
		con: newCon,
	}
}

func (playlistRep *PlaylistRepository) CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {
	query := `INSERT INTO playlists (tittle, description, picture, release_date, user_id) VALUES 
	($1, $2, $3, $4, $5) returning playlist_id`
	err := playlistRep.con.QueryRow(query, playlist.Tittle, playlist.Description,
		playlist.Picture, playlist.ReleaseDate, playlist.UserID).Scan(&playlist.PlaylistID)
	if err != nil {
		return nil, err
	}
	if len(playlist.Tracks) != 0 {
		queryTracksToPlaylists := ``
		for _, track := range playlist.Tracks {
			queryTracksToPlaylists += fmt.Sprintf(
				" INSERT INTO Tracks_to_Playlist (track_id, playlist_id) VALUES (%d, %d);",
				track.TrackID, playlist.PlaylistID)
		}
		_, err = playlistRep.con.Exec(queryTracksToPlaylists)
		if err != nil {
			return nil, err
		}
	}
	queryPlaylistToUser := `INSERT INTO playlists_to_user (user_id, playlist_id) VALUES ($1, $2)`
	_, err = playlistRep.con.Exec(queryPlaylistToUser, playlist.UserID, playlist.PlaylistID)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistRep *PlaylistRepository) DeletePlaylistFromUser(userID, playlistID int) error {
	query := `DELETE FROM playlists_to_user where playlist_id = $1 and user_id = $2`

	_, err := playlistRep.con.Exec(query, playlistID, userID)
	return err
}

func (playlistRep *PlaylistRepository) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	queryGetPlaylist := `SELECT playlist_id, tittle, description, picture, release_date, user_id FROM playlists
						WHERE playlist_id = $1`
	playlist := &models.Playlist{}
	err := playlistRep.con.QueryRow(queryGetPlaylist, playlistID).Scan(
		&playlist.PlaylistID, &playlist.Tittle, &playlist.Description,
		&playlist.Picture, &playlist.ReleaseDate, &playlist.UserID)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}

func (playlistRep *PlaylistRepository) GetTracksByPlaylistID(playlistID int) ([]*trackModels.Track, error) {
	query := `select t.track_id, t.tittle, t.text, t.audio, t.picture, t.release_date from tracks_to_playlist as t_p
			left outer join tracks as t on t.track_id = t_p.track_id
			where playlist_id = $1`
	rows, err := playlistRep.con.Query(query, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tracksByPlaylistID := make([]*trackModels.Track, 0)

	for rows.Next() {
		track := &trackModels.Track{}
		err = rows.Scan(&track.TrackID, &track.Tittle, &track.Text, &track.Audio, &track.Picture,
			&track.ReleaseDate)
		if err != nil {
			return nil, err
		}
		tracksByPlaylistID = append(tracksByPlaylistID, track)
	}
	return tracksByPlaylistID, err
}

func (playlistRep *PlaylistRepository) AddPlaylistToMediateka(userID, playlistID int) error {
	query := `INSERT INTO playlists_to_user (user_id, playlist_id) VALUES ($1, $2)`

	_, err := playlistRep.con.Exec(query, userID, playlistID)
	return err
}

func (plalistRep *PlaylistRepository) GetMediateka(userID int) ([]*models.Playlist, error) {
	query := `select p.playlist_id, p.tittle, p.description, p.picture,
       p.release_date, p.user_id from playlists_to_user as p_u
			left join playlists p on p_u.playlist_id = p.playlist_id
			where p.user_id = $1`

	rows, err := plalistRep.con.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	playlists := make([]*models.Playlist, 0)

	for rows.Next() {
		playlist := &models.Playlist{}
		err = rows.Scan(&playlist.PlaylistID, &playlist.Tittle, &playlist.Description,
			&playlist.Picture, &playlist.ReleaseDate, &playlist.UserID)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, err
}

func (playlistRep *PlaylistRepository) SearchPlaylists(searchQuery string) ([]*models.Playlist, error) {
	query := `SELECT p.playlist_id, p.tittle, p.description, p.picture,
    		p.release_date, p.user_id FROM playlists as p
			WHERE p.tittle LIKE '%' || $1 || '%'`

	rows, err := playlistRep.con.Query(query, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	playlists := make([]*models.Playlist, 0)

	for rows.Next() {
		playlist := &models.Playlist{}
		err = rows.Scan(&playlist.PlaylistID, &playlist.Tittle, &playlist.Description,
			&playlist.Picture, &playlist.ReleaseDate, &playlist.UserID)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (playlistRep *PlaylistRepository) GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error) {
	query := `select p.playlist_id, p.tittle, p.description, p.picture, p.release_date,
     		p.user_id from playlists_to_genres as p_g
			left join playlists p on p.playlist_id = p_g.playlist_id
			where p_g.genre_id = $1`
	rows, err := playlistRep.con.Query(query, genreID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	playlists := make([]*models.Playlist, 0)

	for rows.Next() {
		playlist := &models.Playlist{}
		err = rows.Scan(&playlist.PlaylistID, &playlist.Tittle, &playlist.Description,
			&playlist.Picture, &playlist.ReleaseDate, &playlist.UserID)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (playlistRep *PlaylistRepository) UploadPicture(playlistID int, audioPath string) error {
	query := `UPDATE playlists SET picture = $1
			WHERE playlist_id = $2`
	_, err := playlistRep.con.Exec(query, audioPath, playlistID)
	if err != nil {
		return err
	}
	return nil
}

func (playlistRep *PlaylistRepository) GetPlaylists() ([]*models.Playlist, error) {
	query := `select playlist_id, tittle, description, picture, release_date, user_id
			from playlists
			order by playlist_id`
	rows, err := playlistRep.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	playlists := make([]*models.Playlist, 0)

	for rows.Next() {
		playlist := &models.Playlist{}
		err = rows.Scan(&playlist.PlaylistID, &playlist.Tittle, &playlist.Description,
			&playlist.Picture, &playlist.ReleaseDate, &playlist.UserID)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}