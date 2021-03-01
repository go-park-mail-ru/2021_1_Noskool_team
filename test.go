package main

import (
	"2021_1_Noskool_team/internal/app/tracks/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func countUniqStrings(data *[]string) map[string]int {
	fmt.Println(*data)
	keys := make(map[string]int)
	for _, entry := range *data {
		if _, ok := keys[entry]; !ok {
			keys[entry] = 1
		} else {
			keys[entry]++
		}
	}
	return keys
}

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 dbname=music_service sslmode=disable",
	)
	if err != nil {
		fmt.Println(err)
	}
	tracksRep := repository.NewTracksRepository(db)

	tracks, err := tracksRep.GetTrackByMusicianID(2)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range tracks {
		fmt.Println(*item)
	}
}
