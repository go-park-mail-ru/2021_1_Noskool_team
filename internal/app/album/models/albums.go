package models

type Album struct {
	AlbumID     int    `json:"-"`
	Tittle      string `json:"tittle"`
	Picture     string `json:"picture"`
	ReleaseDate string `json:"release_date"`
}
