package models

//easyjson:json
type Musician struct {
	MusicianID  int    `json:"musician_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

//easyjson:json
type Musicians []*Musician
