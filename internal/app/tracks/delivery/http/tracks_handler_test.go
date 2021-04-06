package http

import (
	"2021_1_Noskool_team/configs"
	mock_tracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	"encoding/json"
	"errors"
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

func TestGetTrackByIDHandlerFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTrackByID(testTreks[0].TrackID).Times(1).Return(testTreks[0],
		errors.New("tracksUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(testTreks[0].TrackID)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTrackByIDHandler(w, r)

	expected := http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

	handler.GetTracksByMusicinID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetTrackByMusicianIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTrackByMusicianID(1).Return(testTreks,
		errors.New("tracksUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(1)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTracksByMusicinID(w, r)

	expected := http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

func TestGetTracksByTittleFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTracksByTittle("tittle").Return(testTreks,
		errors.New("tracksUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_tittle": "tittle"})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTracksByTittle(w, r)

	expected := http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetTracksByGenreID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTracksByGenreID(1).Return(testTreks, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/genre/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre_id": strconv.Itoa(1)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTracksByGenreIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecaseFailed := mock_tracks.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/genre/", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecaseFailed)
	handler.GetTracksByGenreIDHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockTracksUsecaseEmptyBody := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecaseEmptyBody.EXPECT().GetTracksByGenreID(1).Return(nil,
		errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/genre/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre_id": strconv.Itoa(1)})
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecaseEmptyBody)
	handler.GetTracksByGenreIDHandler(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetTracksByAlbumID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)

	mockTracksUsecase.EXPECT().GetTracksByAlbumID(1).Return(testTreks, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(1)})

	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTracksByAlbumIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecaseFailed := mock_tracks.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/album/", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecaseFailed)
	handler.GetTracksByAlbumIDHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockTracksUsecaseEmptyBody := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecaseEmptyBody.EXPECT().GetTracksByAlbumID(1).Return(nil,
		errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(1)})
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecaseEmptyBody)
	handler.GetTracksByAlbumIDHandler(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}
