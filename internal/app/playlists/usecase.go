package playlists

import "2021_1_Noskool_team/internal/app/playlists/models"

type Usecase interface {
	CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error)
	DeletePlaylistFromUser(userID, playlistID int) error
	GetPlaylistByID(playlistID int) (*models.Playlist, error)
	AddPlaylistToMediateka(userID, playlistID int) error
	GetMediateka(userID int) ([]*models.Playlist, error)
	GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error)
	UploadPicture(playlistID int, audioPath string) error
	GetPlaylists() ([]*models.Playlist, error)
	AddTrackToPlaylist(playlistID, trackID int) error
	DeleteTrackFromPlaylist(playlistID, trackID int) error
	UpdatePlaylistTittle(playlist *models.Playlist) error
	UpdatePlaylistDescription(playlist *models.Playlist) error
	GetPlaylistByUID(UID string) (*models.Playlist, error)
}
