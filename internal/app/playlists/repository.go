package playlists

import (
	"2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
)

type Repository interface {
	CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error)
	DeletePlaylistFromUser(userID, playlistID int) error
	GetTracksByPlaylistID(playlistID int) ([]*trackModels.Track, error)
	GetPlaylistByID(playlistID int) (*models.Playlist, error)
	AddPlaylistToMediateka(userID, playlistID int) error
	GetMediateka(userID int) ([]*models.Playlist, error)
	SearchPlaylists(searchQuery string) ([]*models.Playlist, error)
	GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error)
	UploadPicture(playlistID int, audioPath string) error
	GetPlaylists() ([]*models.Playlist, error)
	AddTrackToPlaylist(playlistID, trackID int) error
	DeleteTrackFromPlaylist(playlistID, trackID int) error
	UpdatePlaylistTittle(playlist *models.Playlist) error
	UpdatePlaylistDescription(playlist *models.Playlist) error
	GetPlaylistByUID(UID string) (*models.Playlist, error)
}
