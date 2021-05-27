package http

import (
	"2021_1_Noskool_team/configs"
	mock_album "2021_1_Noskool_team/internal/app/album/mocks"
	"2021_1_Noskool_team/internal/app/album/models"
	mock_musicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	mock_tracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	models2 "2021_1_Noskool_team/internal/microservices/auth/models"
	commonModels "2021_1_Noskool_team/internal/models"
	"context"
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

	albumsForTests = []*models.Album{
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
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum, nil)
	mockTracksUsecase.EXPECT().GetTracksByAlbumID(gomock.Any()).AnyTimes()
	mockMusicianUsecase.EXPECT().GetMusicianByAlbumID(gomock.Any()).AnyTimes()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.GetAlbumByID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	albumWithExtra := ConvertAlumToFullAlbum(testAlbum)
	expectedMsg, _ := json.Marshal(albumWithExtra)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbums().Return(albumsForTests, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.GetAlbums(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(albumsForTests)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestGetAlbumByIDFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumByID(testAlbum.AlbumID).Times(1).Return(testAlbum,
		errors.New("albumUsecase error"))
	mockTracksUsecase.EXPECT().GetTracksByAlbumID(gomock.Any()).AnyTimes()
	mockMusicianUsecase.EXPECT().GetMusicianByAlbumID(gomock.Any()).AnyTimes()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/album/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": strconv.Itoa(testAlbum.AlbumID)})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

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
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByMusicianID(1).Times(1).Return(expectedAlbums, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

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
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByMusicianID(1).Times(1).Return(expectedAlbums,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

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
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByTrackID(1).Times(1).Return(expectedAlbums, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

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
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().GetAlbumsByTrackID(1).Times(1).Return(expectedAlbums,
		errors.New("albumUsecase error"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bytrack/", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": "1"})

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.GetAlbumsByTrackID(w, r)

	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	if !reflect.DeepEqual("{\"status\":\"failed\"}", w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", "{\"status\":\"failed\"}", w.Body.String())
	}
}

func TestAddDeleteAlbumToMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().AddAlbumToMediateka(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r = mux.SetURLVars(r, map[string]string{"album_id": "1"})
	r.URL.RawQuery = "type=add"

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.AddDeleteAlbumToMediateka(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToMediateka(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "cont correct id"})) //nolint
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToMediateka(w, r)

	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "cont correct id"}))
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToMediateka(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestAddDeleteAlbumToFavorites(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockAlbumUsecase.EXPECT().AddAlbumToFavorites(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r = mux.SetURLVars(r, map[string]string{"album_id": "1"})
	r.URL.RawQuery = "type=add"

	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.AddDeleteAlbumToFavorites(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToFavorites(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "cont correct id"})) //nolint
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToFavorites(w, r)

	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "cont correct id"}))
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.AddDeleteAlbumToFavorites(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetFavoriteAlbums(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockMusicianUsecase := mock_musicians.NewMockUsecase(ctrl)
	pagination := &commonModels.Pagination{Limit: 1, Offset: 1}
	mockAlbumUsecase.EXPECT().GetFavoriteAlbums(1, pagination).Return(albumsForTests, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "limit=1&offset=1"
	handler := NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)

	handler.GetFavoriteAlbums(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(albumsForTests)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.GetFavoriteAlbums(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/album/bymusician/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "cont correct id"})) //nolint
	handler = NewAlbumsHandler(mux.NewRouter(), configs.NewConfig(), mockAlbumUsecase,
		mockTracksUsecase, mockMusicianUsecase)
	handler.GetFavoriteAlbums(w, r)

	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}
