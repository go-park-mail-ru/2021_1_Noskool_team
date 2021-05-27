package http

import (
	"2021_1_Noskool_team/configs"
	mock_musicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
	models2 "2021_1_Noskool_team/internal/microservices/auth/models"
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
	testMusicians1 = []*models.Musician{
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
	mockMusiciansUsecase.EXPECT().CheckMusicianInFavorite(1, 1).Return(nil)
	mockMusiciansUsecase.EXPECT().CheckMusicianInMediateka(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMusicianByID(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
	newMusician := models.MusicianFullInformation{
		MusicianID:  testMusicians[0].MusicianID,
		Name:        testMusicians[0].Name,
		Description: testMusicians[0].Description,
		Picture:     testMusicians[0].Picture,
		InMediateka: true,
		InFavorite:  true,
	}

	expectedMsg, _ := json.Marshal(newMusician)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetMusicianByID(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "error id"})) //nolint

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetMusicianByID(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetMediatekaForUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansMediateka(testMusicians[0].MusicianID).Times(1).Return(testMusicians1, nil)
	mockMusiciansUsecase.EXPECT().CheckMusicianInFavorite(gomock.Any(), gomock.Any()).Return(fmt.Errorf("some err")).AnyTimes()
	mockMusiciansUsecase.EXPECT().CheckMusicianInMediateka(gomock.Any(), gomock.Any()).Return(fmt.Errorf("some err")).AnyTimes()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMediatekaForUser(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	musiciansFullInf := make([]*models.MusicianFullInformation, 0)

	for _, item := range testMusicians1 {
		newMusician := &models.MusicianFullInformation{
			MusicianID:  item.MusicianID,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			InMediateka: false,
			InFavorite:  false,
		}
		musiciansFullInf = append(musiciansFullInf, newMusician)
	}

	expectedMsg, _ := json.Marshal(musiciansFullInf)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetMediatekaForUser(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "error id"})) //nolint

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetMediatekaForUser(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestFavoritesForUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusiciansFavorites(testMusicians[0].MusicianID).Times(1).Return(testMusicians1, nil)
	mockMusiciansUsecase.EXPECT().CheckMusicianInFavorite(gomock.Any(), gomock.Any()).Return(fmt.Errorf("some err")).AnyTimes()
	mockMusiciansUsecase.EXPECT().CheckMusicianInMediateka(gomock.Any(), gomock.Any()).Return(fmt.Errorf("some err")).AnyTimes()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetFavoritesForUser(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	musiciansFullInf := make([]*models.MusicianFullInformation, 0)

	for _, item := range testMusicians1 {
		newMusician := &models.MusicianFullInformation{
			MusicianID:  item.MusicianID,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			InMediateka: false,
			InFavorite:  false,
		}
		musiciansFullInf = append(musiciansFullInf, newMusician)
	}

	expectedMsg, _ := json.Marshal(musiciansFullInf)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetFavoritesForUser(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "error id"})) //nolint

	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.GetFavoritesForUser(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestGetMusician(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicians().Times(1).Return(&testMusicians, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMusicians(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians)
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
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

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

func TestAddDeleteMusiciansToMediateka(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)
	mockMusiciansUsecase.EXPECT().AddMusicianToMediateka(1, 1).Times(1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/1/favorite", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "type=add"
	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.AddDeleteMusicianToMediateka(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToMediateka(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToMediateka(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "1"}))
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToMediateka(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestAddDeleteMusiciansToFavorite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)
	mockMusiciansUsecase.EXPECT().AddMusicianToFavorites(1, 1).Times(1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/1/favorite", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(1)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint
	r.URL.RawQuery = "type=add"
	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.AddDeleteMusicianToFavorites(w, r)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToFavorites(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/track/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "not correct id"}))
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToFavorites(w, r)
	expected = http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/vi/musician/1/favorite", nil)
	r = r.WithContext(context.WithValue(r.Context(), "user_id", //nolint
		models2.Result{ID: "1"}))
	handler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler.AddDeleteMusicianToFavorites(w, r)
	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/musicians/byalbum/", nil)
	MusicHandler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler = http.HandlerFunc(MusicHandler.GetMusicianByAlbumID)
	handler.ServeHTTP(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/musicians/byalbum/", nil)
	r = mux.SetURLVars(r, map[string]string{"album_id": "error id"})
	MusicHandler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler = http.HandlerFunc(MusicHandler.GetMusicianByAlbumID)
	handler.ServeHTTP(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/musicians/byplaylist/", nil)
	MusicHandler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	handler = http.HandlerFunc(MusicHandler.GetMusicianByPlaylistID)
	handler.ServeHTTP(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/api/v1/musicians/byplaylist/", nil)
	MusicHandler = NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)
	r = mux.SetURLVars(r, map[string]string{"playlist_id": "error id"})
	handler = http.HandlerFunc(MusicHandler.GetMusicianByPlaylistID)
	handler.ServeHTTP(w, r)

	expected = http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
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
