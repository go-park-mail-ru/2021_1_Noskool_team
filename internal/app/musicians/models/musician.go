package models

import (
	"encoding/json"
	"errors"
)

//easyjson:json
type Musician struct {
	MusicianID  int    `json:"musician_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

//easyjson:json
type MusicianFullInformation struct {
	MusicianID  int    `json:"musician_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	InMediateka bool   `json:"in_mediateka"`
	InFavorite  bool   `json:"in_favorite"`
}

//easyjson:json
type Musicians []*Musician

func MarshalMusician(data interface{}) ([]byte, error) {
	track, ok := data.(*Musician)
	if !ok {
		return nil, errors.New("cant convernt interface{} to musician")
	}
	body, err := track.MarshalJSON()
	return body, err
}

func MarshalMusicianFullInform(data interface{}) ([]byte, error) {
	track, ok := data.(*MusicianFullInformation)
	if !ok {
		return nil, errors.New("cant convernt interface{} to musician")
	}
	body, err := track.MarshalJSON()
	return body, err
}

func MarshalMusicians(data interface{}) ([]byte, error) {
	//track, ok := data.(Musicians)
	//if ok != false {
	//	return nil, errors.New("cant convernt interface{} to musicians")
	//}
	//body, err := track.MarshalJSON()
	body, err := json.Marshal(data)
	return body, err
}
