package playlists

import "2021_1_Noskool_team/internal/app/playlists/models"

type Usecase interface {
	CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error)
	DeletePlaylistFromUser(userID, playlistID int) error
	GetPlaylistByID(playlistID int) (*models.Playlist, error)
	AddPlaylistToMediateka(userID, playlistID int) error
	GetMediateka(userID int) ([]*models.Playlist, error)
}
