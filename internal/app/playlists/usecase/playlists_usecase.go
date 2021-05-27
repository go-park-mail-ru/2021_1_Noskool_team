package usecase

import (
	"2021_1_Noskool_team/internal/app/playlists"
	"2021_1_Noskool_team/internal/app/playlists/models"
	"fmt"
)

type PlaylistUsecase struct {
	playlistRep playlists.Repository
}

func NewPlaylistUsecase(newRep playlists.Repository) *PlaylistUsecase {
	return &PlaylistUsecase{
		playlistRep: newRep,
	}
}

func (usecase *PlaylistUsecase) GetPlaylistByUID(UID string) (*models.Playlist, error) {
	playlist, err := usecase.playlistRep.GetPlaylistByUID(UID)
	if err != nil {
		return nil, err
	}
	tracks, err := usecase.playlistRep.GetTracksByPlaylistID(playlist.PlaylistID)
	if err != nil {
		return nil, err
	}
	playlist.Tracks = tracks
	return playlist, err
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

func (usecase *PlaylistUsecase) GetPlaylistsByGenreID(genreID int) ([]*models.Playlist, error) {
	playlist, err := usecase.playlistRep.GetPlaylistsByGenreID(genreID)
	return playlist, err
}

func (usecase *PlaylistUsecase) UploadPicture(playlistID int, audioPath string) error {
	err := usecase.playlistRep.UploadPicture(playlistID, audioPath)
	return err
}

func (usecase *PlaylistUsecase) GetPlaylists() ([]*models.Playlist, error) {
	playlists, err := usecase.playlistRep.GetPlaylists()
	for _, playlist := range playlists {
		fmt.Println(*playlist)
		tracks, _ := usecase.playlistRep.GetTracksByPlaylistID(playlist.PlaylistID)
		playlist.Tracks = tracks
	}
	return playlists, err
}

func (usecase *PlaylistUsecase) AddTrackToPlaylist(playlistID, trackID int) error {
	err := usecase.playlistRep.AddTrackToPlaylist(playlistID, trackID)
	return err
}

func (usecase *PlaylistUsecase) DeleteTrackFromPlaylist(playlistID, trackID int) error {
	err := usecase.playlistRep.DeleteTrackFromPlaylist(playlistID, trackID)
	return err
}

func (usecase *PlaylistUsecase) UpdatePlaylistTittle(playlist *models.Playlist) error {
	err := usecase.playlistRep.UpdatePlaylistTittle(playlist)
	return err
}

func (usecase *PlaylistUsecase) UpdatePlaylistDescription(playlist *models.Playlist) error {
	err := usecase.playlistRep.UpdatePlaylistDescription(playlist)
	return err
}
