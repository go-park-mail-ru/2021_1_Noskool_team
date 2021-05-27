package utility

import (
	"2021_1_Noskool_team/configs"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	if err == nil {
		defer con.Close()
		t.Error(err)
	}
}

func TestCheckUserIDFail1(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"user_id": "1"})

	logg := logrus.New()
	_, err := CheckUserID(w, r, logg)

	assert.Equal(t, err, errors.New("Not correct user id"))
}
