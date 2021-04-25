package models

import "2021_1_Noskool_team/internal/app/tracks/models"

//easyjson:json
type Playlist struct {
	PlaylistID  int             `json:"playlist_id"`
	Tittle      string          `json:"tittle"`
	Description string          `json:"description"`
	Picture     string          `json:"picture"`
	ReleaseDate string          `json:"release_date"`
	UserID      int             `json:"user_id"`
	Tracks      []*models.Track `json:"tracks"`
}

//easyjson:json
type Playlists []*Playlist