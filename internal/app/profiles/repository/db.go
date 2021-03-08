package repository

import (
	"2021_1_Noskool_team/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config            *configs.Config
	con               *sql.DB
	profileRepository *ProfileRepository
}

// New ...
func New(config *configs.Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	fmt.Println(s.config.ProfileDB)
	db, err := sql.Open("postgres", "host=localhost dbname=profiles_db user=server password=password sslmode=disable") // s.config.ProfileDB "host=localhost dbname=profiles_db user=server password=password sslmode=disable"
	if err != nil {
		fmt.Println("open db")
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("ping db")
		return err
	}

	s.con = db
	return nil
}

// Close ...
func (s *Store) Close() {
	s.con.Close()
}

// User ...
func (s *Store) User() *ProfileRepository {
	if s.profileRepository != nil {
		return s.profileRepository
	}

	s.profileRepository = &ProfileRepository{
		db: s,
	}
	return s.profileRepository
}
