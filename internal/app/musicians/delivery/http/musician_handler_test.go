package http

import (
	"2021_1_Noskool_team/configs"
	mock_musicians "2021_1_Noskool_team/internal/app/musicians/mocks"
	"2021_1_Noskool_team/internal/app/musicians/models"
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
	testMusicians = []*models.Musician{
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

func TestGetMusicByIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockMusiciansUsecase := mock_musicians.NewMockUsecase(ctrl)

	mockMusiciansUsecase.EXPECT().GetMusicianByID(testMusicians[0].MusicianID).Times(1).Return(testMusicians[0], nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/musician/", nil)
	r = mux.SetURLVars(r, map[string]string{"musician_id": strconv.Itoa(testMusicians[0].MusicianID)})

	handler := NewMusicHandler(mux.NewRouter(), configs.NewConfig(), mockMusiciansUsecase)

	handler.GetMusicByIDHandler(w, r)

	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(testMusicians[0])
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}
