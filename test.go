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

	musicians, err := musicRep.GetMusiciansByGenres("pop")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(musicians)
}
