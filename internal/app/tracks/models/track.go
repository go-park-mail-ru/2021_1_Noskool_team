package models

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	"2021_1_Noskool_team/internal/models"
	"errors"
)

//easyjson:json
type Track struct {
	TrackID     int                         `json:"track_id"`
	Tittle      string                      `json:"tittle"`
	Text        string                      `json:"text"`
	Audio       string                      `json:"audio"`
	Picture     string                      `json:"picture"`
	ReleaseDate string                      `json:"release_date"`
	Duration    string                      `json:"duration"`
	InMediateka bool                        `json:"in_mediateka"`
	InFavorite  bool                        `json:"in_favorite"`
	Likes       int                         `json:"likes"`
	Genres      []*models.Genre             `json:"genres"`
	Musicians   []*musiciansModels.Musician `json:"musicians"`
	Albums      []*albumsModels.Album       `json:"album"`
}

//easyjson:json
type Tracks []*Track

func MarshalTracks(data interface{}) ([]byte, error) {
	track, ok := data.([]*Track)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	var tracks Tracks //nolint
	tracks = track
	body, err := tracks.MarshalJSON()
	return body, err
}

func MarshalTrack(data interface{}) ([]byte, error) {
	track, ok := data.(*Track)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := track.MarshalJSON()
	return body, err
}
