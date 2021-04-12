package http

import (
	"2021_1_Noskool_team/configs"
	mock_musicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var (
	testMusicians = []models.Musician{
		{
			MusicianID:  1,
			Name:        "some musician",
			Description: "description",
			Picture:     "picture",
		},
		{
			MusicianID:  2,
			Name:        "some musician2",
			Description: "description2",
			Picture:     "picture2",
		},
	}
)

func TestGetMusicianByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByID(testMusicians[0].MusicianID).Times(1).Return(&testMusicians[0], nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMusicianByID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians[0])
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusicianByIDFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByID(testMusicians[0].MusicianID).Times(1).Return(&testMusicians[0],
		errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMusicianByID(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetMusiciansByGenre(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansByGenre("rok").Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/bygenre/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre": "rok"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusiciansByGenre)
	handler.ServeHTTP(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusiciansByGenreFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansByGenre("rok").Times(1).
		Return(&testMusicians, errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/bygenre/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre": "rok"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusiciansByGenre)
	handler.ServeHTTP(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetMusicianByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByTrackID(1).Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByTrackID)
	handler.ServeHTTP(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusicianByTrackIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByTrackID(1).Times(1).
		Return(&testMusicians, errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByTrackID)
	handler.ServeHTTP(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetMusicianByAlbumID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByAlbumID(1).Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/byalbum/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByAlbumID)
	handler.ServeHTTP(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusicianByAlbumIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByAlbumID(1).Times(1).
		Return(&testMusicians, errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/byalbum/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByAlbumID)
	handler.ServeHTTP(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetMusicianByPlaylistID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByPlaylistID(1).Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/byplaylist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByPlaylistID)
	handler.ServeHTTP(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusicianByPlaylistIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByPlaylistID(1).Times(1).
		Return(&testMusicians, errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/byplaylist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": "1"})

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusicianByPlaylistID)
	handler.ServeHTTP(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetMusiciansTop3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansTop4().Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/poular", nil)

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusiciansTop4)
	handler.ServeHTTP(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetMusiciansTop3Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansTop4().Times(1).
		Return(&testMusicians, errors.New("new err"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/musicians/poular", nil)

	MusicHandler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler := http.HandlerFunc(MusicHandler.GetMusiciansTop4)
	handler.ServeHTTP(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}
