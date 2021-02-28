package models

type Album struct {
	AlbumID     int    `json:"album_id"`
	Tittle      string `json:"tittle"`
	Picture     string `json:"picture"`
	ReleaseDate string `json:"release_date"`
}
