package http

import (
	"2021_1_Noskool_team/configs"
	mock_tracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	"2021_1_Noskool_team/internal/app/tracks/models"
	models2 "2021_1_Noskool_team/internal/microservices/auth/models"
	models0 "2021_1_Noskool_team/internal/models"
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

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTrackByIDHandler(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/musician/", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetTracksByMusicinID(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

func TestAddDeleteTrackToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().AddTrackToFavorites(1, 1).Return(nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "type=add"
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.AddDeleteTrackToFavorite(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToFavorite(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToFavorite(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "1"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToFavorite(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestAddDeleteTrackToMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().AddDeleteTrackToMediateka(1, 1, "add")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/1/mediateka", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "type=add"
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.AddDeleteTrackToMediateka(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/mediateka", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToMediateka(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/mediateka", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToMediateka(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "1"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddDeleteTrackToMediateka(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockTracksUsecase.EXPECT().AddDeleteTrackToMediateka(1, 1,
		"add").Return(errors.New("no content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/mediateka", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "type=add"
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.AddDeleteTrackToMediateka(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetFavoriteTracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetFavoriteTracks(1, &models0.Pagination{
		Limit:  10,
		Offset: 0,
	}).Return(testTreks, nil)
	mockTracksUsecase.EXPECT().CheckTrackInMediateka(gomock.Any(), gomock.Any()).AnyTimes()
	mockTracksUsecase.EXPECT().CheckTrackInFavorite(gomock.Any(), gomock.Any()).AnyTimes()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/favorites", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "limit=10&offset=0"
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetFavoriteTracks(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/favorites", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetFavoriteTracks(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/favorites", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetFavoriteTracks(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockTracksUsecase.EXPECT().GetFavoriteTracks(1, &models0.Pagination{
		Limit:  10,
		Offset: 0,
	}).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/favorites", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "limit=10&offset=0"
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetFavoriteTracks(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetMediatekaForUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetTracksByUserID(1).Return(testTreks, nil)
	mockTracksUsecase.EXPECT().CheckTrackInMediateka(gomock.Any(), gomock.Any()).AnyTimes()
	mockTracksUsecase.EXPECT().CheckTrackInFavorite(gomock.Any(), gomock.Any()).AnyTimes()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/mediateka", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetMediatekaForUser(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/mediateka", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetMediatekaForUser(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/favorites", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetMediatekaForUser(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockTracksUsecase.EXPECT().GetTracksByUserID(1).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/mediateka", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)

	handler.GetMediatekaForUser(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetHistory(1).Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/history", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetHistory(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/history", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetHistory(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/history", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetHistory(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockTracksUsecase.EXPECT().GetHistory(1).Return(nil, errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/history", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetHistory(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestAddToHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().AddToHistory(1, 1).Return(nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/track/1/history", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddToHistory(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/track/1/history", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddToHistory(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/track/1/history", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddToHistory(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("POST", "/api/vi/track/1/history", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddToHistory(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	mockTracksUsecase.EXPECT().AddToHistory(1, 1).Return(errors.New("some error"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/history", nil)
	r = mux.SetURLVars(r, map[string]string{"track_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.AddToHistory(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetTop20Tracks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetTop20Tracks().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTop20Tracks(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetTop20Tracks().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTop20Tracks(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetTop20TracksNotAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetTop20Tracks().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTop20TracksNotAuth(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetTop20Tracks().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTop20TracksNotAuth(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetTopTrack(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetTopTrack().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTopTrack(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetTopTrack().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTopTrack(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetTopTrackNotAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetTopTrack().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTopTrackNotAuth(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetTopTrack().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/top", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetTopTrackNotAuth(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetBillbordTopCharts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetBillbordTopCharts().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/billbord", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetBillbordTopCharts(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetBillbordTopCharts().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/billbord", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetBillbordTopCharts(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetBillbordTopChartsNotAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockTracksUsecase.EXPECT().GetBillbordTopCharts().Return(testTreks, nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/track/billbord", nil)
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetBillbordTopChartsNotAuth(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testTreks)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	mockTracksUsecase.EXPECT().GetBillbordTopCharts().Return(nil, errors.New("not content"))
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/billbord", nil)
	handler = NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.GetBillbordTopChartsNotAuth(w, r)
	expected = http.StatusNoContent
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestUploadTrackAudioHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	//mockTracksUsecase.EXPECT().UploadAudio(1, "some path").Return(nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/track/billbord", nil)
	r.Header.Set("Content-Type", "multipart/form-data")
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.UploadTrackAudioHandler(w, r)
	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestUploadTrackPictureHandlerr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	//mockTracksUsecase.EXPECT().UploadAudio(1, "some path").Return(nil)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/api/vi/track/billbord", nil)
	r.Header.Set("Content-Type", "multipart/form-data")
	handler := NewTracksHandler(mux.NewRouter(), configs.NewConfig(), mockTracksUsecase)
	handler.UploadTrackPictureHandler(w, r)
	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}
