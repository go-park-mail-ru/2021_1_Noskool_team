package usecase

import (
	"2021_1_Noskool_team/internal/app/playlists"
	"2021_1_Noskool_team/internal/app/playlists/models"
)

type PlaylistUsecase struct {
	playlistRep playlists.Repository
}

func NewPlaylistUsecase(newRep playlists.Repository) *PlaylistUsecase {
	return &PlaylistUsecase{
		playlistRep: newRep,
	}
}

func (usecase *PlaylistUsecase) CreatePlaylist(playlist *models.Playlist) (*models.Playlist, error) {
	playlist, err := usecase.playlistRep.CreatePlaylist(playlist)
	return playlist, err
}

func (usecase *PlaylistUsecase) DeletePlaylistFromUser(userID, playlistID int) error {
	err := usecase.playlistRep.DeletePlaylistFromUser(userID, playlistID)
	return err
}
