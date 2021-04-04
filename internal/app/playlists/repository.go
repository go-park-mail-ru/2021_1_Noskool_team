package playlists

import "2021_1_Noskool_team/internal/app/playlists/models"

type Repository interface {
	CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error)
}
