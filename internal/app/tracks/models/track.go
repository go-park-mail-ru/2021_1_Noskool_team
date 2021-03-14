package models

type Track struct {
	TrackID     int      `json:"track_id"`
	Tittle      string   `json:"tittle"`
	Text        string   `json:"text"`
	Audio       string   `json:"audio"`
	Picture     string   `json:"picture"`
	ReleaseDate string   `json:"release_date"`
	Genres      []string `json:"genres"`
	Musicians   []string `json:"musicians"`
	Albums      []string `json:"album"`
}
