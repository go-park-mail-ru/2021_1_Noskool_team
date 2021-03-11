package http

import (
	"2021_1_Noskool_team/configs"
	mock_album "2021_1_Noskool_team/internal/app/album/mocks"
	"2021_1_Noskool_team/internal/app/album/models"
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
	testAlbum = &models.Album{
		AlbumID:     1,
		Tittle:      "some tittle",
		Picture:     "some picture",
		ReleaseDate: "date",
	}
)

func TestGetAlbumByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumByIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testAlbum)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbumByIDHandlerFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase)

	handler.GetAlbumByIDHandler(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}
