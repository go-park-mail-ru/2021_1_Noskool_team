package usecase

import "2021_1_Noskool_team/internal/app/playlists"

type PlaylistUsecase struct {
	playlistRep playlists.Repository
}

func NewPlaylistUsecase(newRep *playlists.Repository) *PlaylistUsecase {
	return &PlaylistUsecase{
		playlistRep: newRep,
	}
}
