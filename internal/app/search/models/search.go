package models

import (
	albumsModels "2021_1_Noskool_team/internal/app/album/models"
	musiciansModels "2021_1_Noskool_team/internal/app/musicians/models"
	playlistsModels "2021_1_Noskool_team/internal/app/playlists/models"
	trackModels "2021_1_Noskool_team/internal/app/tracks/models"
	"2021_1_Noskool_team/internal/models"
	"errors"
)

//easyjson:json
type Search struct {
	Tracks    []*TrackWithAlbum        `json:"tracks"`
	Albums    []*albumsModels.Album       `json:"albums"`
	Musicians []*musiciansModels.Musician `json:"musicians"`
	Playlists []*playlistsModels.Playlist `json:"playlists"`
}

//easyjson:json
type TrackWithAlbum struct {
	TrackID     int                         `json:"track_id"`
	Tittle      string                      `json:"tittle"`
	Text        string                      `json:"text"`
	Audio       string                      `json:"audio"`
	Picture     string                      `json:"picture"`
	ReleaseDate string                      `json:"release_date"`
	Duration    string                      `json:"duration"`
	InMediateka bool                        `json:"in_mediateka"`
	InFavorite  bool                        `json:"in_favorite"`
	Genres      []*models.Genre             `json:"genres"`
	Musicians   []*musiciansModels.Musician `json:"musicians"`
	Albums      []*albumsModels.Album       `json:"album"`
	Album int
}

func ConvertTrackToTrackWithAlbum(track *trackModels.Track) *TrackWithAlbum {
	return &TrackWithAlbum{
		TrackID:     track.TrackID,
		Tittle:      track.Tittle,
		Text:        track.Text,
		Audio:       track.Audio,
		Picture:     track.Picture,
		ReleaseDate: track.ReleaseDate,
		Duration:    track.Duration,
		InMediateka: track.InMediateka,
		InFavorite:  track.InFavorite,
		Genres:      track.Genres,
		Musicians:   track.Musicians,
		Albums:      track.Albums,
		Album:       0,
	}
}

func MarshalSearch(data interface{}) ([]byte, error) {
	track, ok := data.(*Search)
	if !ok {
		return nil, errors.New("cant convernt interface{} to track")
	}
	body, err := track.MarshalJSON()
	return body, err
}
