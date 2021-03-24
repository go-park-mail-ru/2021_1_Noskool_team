package repository

import "database/sql"

type PlaylistRepository struct {
	con *sql.DB
}

func NewPlaylistRepository(newCon *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{
		con: newCon,
	}
}
