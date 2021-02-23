package main

import (
	"2021_1_Noskool_team/internal/app/music/repository"
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
	musicRep := repository.NewMusicRepository(db)

	track, err := musicRep.GetTrackById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(track)
}
