package main

import (
	"2021_1_Noskool_team/internal/app/tracks/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 dbname=music_service sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
	}
	tracksRep := repository.NewTracksRepository(db)

	tracks, err := tracksRep.GetTracksByTittle("song23")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range tracks {
		fmt.Println(*item)
	}
}
