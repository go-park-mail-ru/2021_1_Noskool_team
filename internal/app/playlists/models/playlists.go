package models

import (
	"2021_1_Noskool_team/internal/app/tracks/models"
	"encoding/json"
	"errors"
)

//easyjson:json
type Playlist struct {
	PlaylistID  int             `json:"playlist_id"`
	Tittle      string          `json:"tittle"`
	Description string          `json:"description"`
	Picture     string          `json:"picture"`
	ReleaseDate string          `json:"release_date"`
	UserID      int             `json:"user_id"`
	Tracks      []*models.Track `json:"tracks"`
	UID         string          `json:"uid"`
}

//easyjson:json
type Playlists []*Playlist

func MarshalPlaylist(data interface{}) ([]byte, error) {
	track, ok := data.(*Playlist)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := track.MarshalJSON()
	return body, err
}

func MarshalPlaylists(data interface{}) ([]byte, error) {
	//track, ok := data.(Playlists)
	//if !ok {
	//	return nil, errors.New("cant convernt interface{} to track")
	//}
	//body, err := track.MarshalJSON()
	body, err := json.Marshal(data)
	return body, err
}
