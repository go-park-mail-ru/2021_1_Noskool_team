package utility

import (
	"2021_1_Noskool_team/configs"
	"testing"
)

func TestCreatePostgresConnectionFailed(t *testing.T) {
	con, err := CreatePostgresConnection("some wrong db settings")
	if con != nil {
		t.Error("Failed")
	}

	if err == nil {
		t.Error("Failed", err)
	}
}

func TestCreatePostgresConnection(t *testing.T) {
	config := configs.NewConfig()
	con, err := CreatePostgresConnection(config.MusicPostgresBD)
	defer con.Close()
	if err != nil {
		t.Error(err)
	}
}
