package repository

import (
	"2021_1_Noskool_team/internal/app/playlists/models"
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
		fmt.Println(queryTracksToPlaylists)
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
