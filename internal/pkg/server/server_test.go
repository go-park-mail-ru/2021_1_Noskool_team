package server

import (
	"2021_1_Noskool_team/configs"
	"github.com/gorilla/mux"
	"testing"
)

func TestFailedResponse(t *testing.T) {
	expected := "{\"status\":\"failed\"}"
	response := string(FailedResponse())

	if response != expected {
		t.Errorf("Failed, expected: %v acctual: %v", expected, response)
	}
}

func TestNewServer(t *testing.T) {
	config := configs.NewConfig()

	serv, err := NewServer(config, mux.NewRouter())
	if err != nil {
		t.Error(err)
	}

	if serv.config.MusicServerAddr != ":8888" {
		t.Errorf("Failed, expected: %v acctual: %v", ":8888", serv.config.MusicServerAddr)
	}
}
