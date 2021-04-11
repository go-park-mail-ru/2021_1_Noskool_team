package http

import (
	"2021_1_Noskool_team/configs"
	mock_album "2021_1_Noskool_team/internal/app/album/mocks"
	"2021_1_Noskool_team/internal/app/album/models"
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
	testAlbum = &models.Album{
		AlbumID:     1,
		Tittle:      "some tittle",
		Picture:     "some picture",
		ReleaseDate: "date",
	}
	expectedAlbums = &[]models.Album{
		{
			AlbumID:     1,
			Tittle:      "album1",
			Picture:     "picture1",
			ReleaseDate: "date1",
		},
		{
			AlbumID:     2,
			Tittle:      "album2",
			Picture:     "picture2",
			ReleaseDate: "date2",
		},
	}
)

func TestGetAlbumByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumByID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testAlbum)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbumByIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumByID(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetAlbumsByMusicianID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByMusicianID(1).Times(1).Return(expectedAlbums, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumsByMusicianID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(expectedAlbums)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbumsByMusicianIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByMusicianID(1).Times(1).Return(expectedAlbums,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumsByMusicianID(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestGetAlbumsByTrackID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByTrackID(1).Times(1).Return(expectedAlbums, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumsByTrackID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(expectedAlbums)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbumsByTrackIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByTrackID(1).Times(1).Return(expectedAlbums,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumsByTrackID(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}
