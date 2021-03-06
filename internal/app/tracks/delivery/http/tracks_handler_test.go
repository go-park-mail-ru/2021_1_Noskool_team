package http

import (
	"2021_1_Noskool_team/configs"
	mock_tracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

var (
	testTreks = []*models.Track{
		{
			TrackID:     1,
			Tittle:      "tittle",
			Text:        "some text",
			Audio:       "audio",
			Picture:     "picture",
			ReleaseDate: "date",
			Genres:      nil,
			Musicians:   nil,
		},
		{
			TrackID:     2,
			Tittle:      "tittle",
			Text:        "some text",
			Audio:       "audio",
			Picture:     "picture",
			ReleaseDate: "date",
			Genres:      nil,
			Musicians:   nil,
		},
	}
)

func TestGetTrackByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTrackByID(testTreks[0].TrackID).Times(1).Return(testTreks[0], nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(testTreks[0].TrackID)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTrackByIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks[0])
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetTrackByMusicianID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTrackByMusicianID(1).Return(testTreks, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(1)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTrackByMusicianID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetTracksByTittle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTracksByTittle("tittle").Return(testTreks, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_tittle": "tittle"})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTracksByTittle(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}
