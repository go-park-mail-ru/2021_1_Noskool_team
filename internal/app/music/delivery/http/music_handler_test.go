package http

import (
	"2021_1_Noskool_team/configs"
	mock_album "2021_1_Noskool_team/internal/app/album/mocks"
	mock_musicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
	mock_playlists "2021_1_Noskool_team/internal/app/playlists/mocks"
	mock_search "2021_1_Noskool_team/internal/app/search/mocks"
	mock_tracks "2021_1_Noskool_team/internal/app/tracks/mocks"
	models2 "2021_1_Noskool_team/internal/microservices/auth/models"
	"context"
	"encoding/json"
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

func TestFinalHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)
	mockTracksUsecase := mock_tracks.NewMockUsecase(ctrl)
	mockAlbumUsecase := mock_album.NewMockUsecase(ctrl)
	mockSearchUsecase := mock_search.NewMockUsecase(ctrl)
	mockPlaylistsUsecase := mock_playlists.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByID(testMusicians[0].MusicianID).Times(1).Return(&testMusicians[0], nil)
	mockMusiciansUsecase.EXPECT().CheckMusicianInMediateka(1, 1).Return(nil)
	mockMusiciansUsecase.EXPECT().CheckMusicianInFavorite(1, 1).Return(nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})
	r = r.WithContext(context.WithValue(r.Context(), "user_id", models2.Result{ID: "1"})) //nolint

	handler := NewFinalHandler(configs.NewConfig(), mockTracksUsecase, mockMusiciansUsecase,
		mockAlbumUsecase, mockPlaylistsUsecase, mockSearchUsecase)

	handler.musicianHandler.GetMusicianByID(w, r)

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
}
