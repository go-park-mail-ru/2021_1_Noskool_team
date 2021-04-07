package models

type Musician struct {
	MusicianID  int    `json:"musician_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`

}
