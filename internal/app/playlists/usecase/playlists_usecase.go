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

func (usecase *PlaylistUsecase) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	playlist, err := usecase.playlistRep.GetPlaylistByID(playlistID)
	if err != nil {
		return nil, err
	}
	tracks, err := usecase.playlistRep.GetTracksByPlaylistID(playlistID)
	if err != nil {
		return nil, err
	}
	playlist.Tracks = tracks
	return playlist, nil
}

func (usecase *PlaylistUsecase) AddPlaylistToMediateka(userID, playlistID int) error {
	err := usecase.playlistRep.AddPlaylistToMediateka(userID, playlistID)
	return err
}

func (usecase *PlaylistUsecase) GetMediateka(userID int) ([]*models.Playlist, error) {
	playlists, err := usecase.playlistRep.GetMediateka(userID)
	if err != nil {
		return nil, err
	}
	for _, playlist := range playlists {
		tracks, err := usecase.playlistRep.GetTracksByPlaylistID(playlist.PlaylistID)
		if err != nil {
			return nil, err
		}
		playlist.Tracks = tracks
	}
	return playlists, nil
}

func (usecase *PlaylistUsecase) GetTracksByGenreID(genreID int) ([]*models.Playlist, error) {
	playlist, err := usecase.playlistRep.GetTracksByGenreID(genreID)
	return playlist, err
}
