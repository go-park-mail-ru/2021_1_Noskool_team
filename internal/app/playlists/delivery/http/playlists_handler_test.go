package http

import (
	"2021_1_Noskool_team/configs"
	mock_album "2021_1_Noskool_team/internal/app/album/mocks"
	mock_playlists "2021_1_Noskool_team/internal/app/playlists/mocks"
	"2021_1_Noskool_team/internal/app/playlists/models"
	models2 "2021_1_Noskool_team/internal/microservices/auth/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var (
	playlistsForTest = []*models.Playlist{
		{
			PlaylistID:  1,
			Tittle:      "Tittle of first playlist",
			Description: "some description",
			Picture:     "/api/v1/data/img/playlists/1.png",
			ReleaseDate: "2020-03-04",
			UserID:      1,
		},
		{
			PlaylistID:  2,
			Tittle:      "Tittle of second playlist",
			Description: "some other description",
			Picture:     "/api/v1/data/img/playlists/1.png",
			ReleaseDate: "2020-07-07",
			UserID:      5,
		},
		{
			PlaylistID:  3,
			Tittle:      "Tittle without tracks",
			Description: "some description",
			Picture:     "/api/v1/data/img/playlists/3.png",
			ReleaseDate: "2020-03-04",
			UserID:      1,
		},
	}
)

func TestDeletePlaylistFromMediatekaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().DeletePlaylistFromUser(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(),
		mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeletePlaylistFromMediatekaHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("DELETE", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeletePlaylistFromMediatekaHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("DELETE", "/api/vi/playlist", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeletePlaylistFromMediatekaHandler(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("DELETE", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeletePlaylistFromMediatekaHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().DeletePlaylistFromUser(1, 1).Return(
		errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("DELETE", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeletePlaylistFromMediatekaHandler(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylistByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylistByID(1).Return(playlistsForTest[0], nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistByIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(playlistsForTest[0])
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistByIDHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().GetPlaylistByID(1).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistByIDHandler(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylistByUIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylistByUID("fdsfds").Return(playlistsForTest[0], nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/fdsfds/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_uid": "fdsfds"})

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistByUIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(playlistsForTest[0])
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase.EXPECT().GetPlaylistByUID("fdsfds").Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_uid": "fdsfds"})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistByUIDHandler(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylistsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetMediateka(1).Return(playlistsForTest, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/fdsfds/", nil)
	r = mux.SetURLVars(r, map[string]string{"user_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByUserID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(playlistsForTest)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_uid": "fdsfds"})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByUserID(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().GetMediateka(1).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"user_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByUserID(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestAddPlaylistFromMediatekaHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().AddPlaylistToMediateka(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddPlaylistToMediatekaHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddPlaylistToMediatekaHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddPlaylistToMediatekaHandler(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddPlaylistToMediatekaHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().AddPlaylistToMediateka(1, 1).Return(
		errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddPlaylistToMediatekaHandler(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylistsByGenreID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylistsByGenreID(1).Return(playlistsForTest, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/genre/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre_id": strconv.Itoa(1)})
	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByGenreID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(playlistsForTest)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/genre/1/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByGenreID(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().GetPlaylistsByGenreID(1).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/genre/1/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsByGenreID(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetMediateka(1).Return(playlistsForTest, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"genre_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetMediateka(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(playlistsForTest)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetMediateka(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "some bad id"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetMediateka(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase.EXPECT().GetMediateka(1).Return(
		nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/playlist/", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetMediateka(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestAddTrackToPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().AddTrackToPlaylist(1, 5).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1),
		"track_id": strconv.Itoa(5)})

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddTrackToPlaylist(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddTrackToPlaylist(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddTrackToPlaylist(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().AddTrackToPlaylist(1,
		5).Return(fmt.Errorf("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1),
		"track_id": strconv.Itoa(5)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.AddTrackToPlaylist(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestDeleteTrackFromPlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().DeleteTrackFromPlaylist(1, 5).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1),
		"track_id": strconv.Itoa(5)})

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeleteTrackFromPlaylist(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeleteTrackFromPlaylist(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeleteTrackFromPlaylist(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().DeleteTrackFromPlaylist(1,
		5).Return(fmt.Errorf("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": strconv.Itoa(1),
		"track_id": strconv.Itoa(5)})
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.DeleteTrackFromPlaylist(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylists().Return(playlistsForTest, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/genre/1/", nil)
	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylists(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(playlistsForTest)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylists().Return(nil, fmt.Errorf("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylists(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestGetPlaylistsNotAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylists().Return(playlistsForTest, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/playlist/genre/1/", nil)
	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsNotAuth(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	expectedMsg, _ := json.Marshal(playlistsForTest)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().GetPlaylists().Return(nil, fmt.Errorf("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.GetPlaylistsNotAuth(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}

func TestCreatePlaylistHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlaylistUsecase := mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().CreatePlaylist(playlistsForTest[0]).Return(playlistsForTest[0], nil)
	body, _ := json.Marshal(playlistsForTest[0])

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/playlist/", bytes.NewBuffer(body))
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.CreatePlaylistHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", nil)
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.CreatePlaylistHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", bytes.NewBuffer(body))
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.CreatePlaylistHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", bytes.NewBuffer(body))
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "not correct id"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.CreatePlaylistHandler(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}

	mockPlaylistUsecase = mock_playlists.NewMockUsecase(ctrl)
	mockAlbumUsecase = mock_album.NewMockUsecase(ctrl)

	mockPlaylistUsecase.EXPECT().CreatePlaylist(playlistsForTest[0]).Return(nil, fmt.Errorf("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/playlist/", bytes.NewBuffer(body))
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewPlaylistsHandler(mux.NewRouter(), configs.NewConfig(), mockPlaylistUsecase, mockAlbumUsecase)
	handler.CreatePlaylistHandler(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("expected: %v\n got: %v", http.StatusBadRequest, w.Code)
	}
}
