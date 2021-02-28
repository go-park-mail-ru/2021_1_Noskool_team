package repository

import (
	"2021_1_Noskool_team/internal/app/album"
	"database/sql"
)

type AlbumsRepository struct {
	con *sql.DB
}

func NewAlbumsRepository(con *sql.DB) album.Repository {
	return &AlbumsRepository{
		con: con,
	}
}
