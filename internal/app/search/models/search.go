package models

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	playlistsModels "2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"errors"
)

//easyjson:json
type Search struct {
	Tracks    []*trackModels.Track        `json:"tracks"`
	Albums    []*albumsModels.Album       `json:"albums"`
	Musicians []*musiciansModels.Musician `json:"musicians"`
	Playlists []*playlistsModels.Playlist `json:"playlists"`
}

func MarshalSearch(data interface{}) ([]byte, error) {
	track, ok := data.(*Search)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := track.MarshalJSON()
	return body, err
}
