package models

import (
	"encoding/json"
	"errors"
)

//easyjson:json
type Album struct {
	AlbumID     int    `json:"album_id"`
	Tittle      string `json:"tittle"`
	Picture     string `json:"picture"`
	ReleaseDate string `json:"release_date"`
}

//easyjson:json
type AlbumFullInformation struct {
	AlbumID     int    `json:"album_id"`
	Tittle      string `json:"tittle"`
	Picture     string `json:"picture"`
	ReleaseDate string `json:"release_date"`
	InMediateka bool   `json:"in_mediateka"`
	InFavorite  bool   `json:"in_favorite"`
}

//easyjson:json
type Albums []*Album

func MarshalAlbum(data interface{}) ([]byte, error) {
	album, ok := data.(*Album)
	if !ok {
		return nil, errors.New("cant convernt interface{} to album")
	}
	body, err := album.MarshalJSON()
	return body, err
}

func MarshalAlbums(data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	return body, err
}
