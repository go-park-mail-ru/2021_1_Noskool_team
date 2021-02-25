package models

type Musician struct {
	MusicianID  int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}
