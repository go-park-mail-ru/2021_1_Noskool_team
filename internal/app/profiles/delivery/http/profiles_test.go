package profiles

import (
	"2021_1_Noskool_team/configs"
	profileModels "2021_1_Noskool_team/internal/app/profiles/models"
	mock_client "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client/mocks"
	"2021_1_Noskool_team/internal/microservices/auth/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func TestHandleAuthWithCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSeshClient := mock_client.NewMockAuthCheckerClient(ctrl)

	session := models.Sessions{
		UserID:     "1",
		Hash:       "fdsjfkdsjelfdlksjfkjds",
		Expiration: 86400,
	}

	result := models.Result{
		ID:     "1",
		Hash:   "fdsjfkdsjelfdlksjfkjds",
		Status: "OK",
	}

	mockSeshClient.EXPECT().Check(context.Background(), session.Hash).Return(result, nil)

	w := httptest.NewRecorder()
	recorder := httptest.NewRecorder()

	http.SetCookie(recorder, &http.Cookie{Name: "session_id", Value: "fdsjfkdsjelfdlksjfkjds"})
	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	handler := ProfilesServer{sessionsClient: mockSeshClient, logger: logrus.New()}

	handler.HandleAuth(w, request)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestHandleAuthNoCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSeshClient := mock_client.NewMockAuthCheckerClient(ctrl)

	session := models.Sessions{
		UserID:     "1",
		Hash:       "wrong cookie",
		Expiration: 86400,
	}

	result := models.Result{
		ID:     "1",
		Hash:   "fdsjfkdsjelfdlksjfkjds",
		Status: "OK",
	}

	mockSeshClient.EXPECT().Check(context.Background(), session.Hash).Return(result, nil)

	w := httptest.NewRecorder()
	recorder := httptest.NewRecorder()

	http.SetCookie(recorder, &http.Cookie{Name: "session_id", Value: "wrong cookie"})
	request := &http.Request{Header: http.Header{"Cookie": recorder.HeaderMap["Set-Cookie"]}}

	handler := ProfilesServer{
		sessionsClient: mockSeshClient,
		router:         mux.NewRouter(),
		logger:         logrus.New(),
		config:         configs.NewConfig(),
	}
	handler.configureLogger()
	handler.configureRouter()

	handler.HandleAuth(w, request)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}
}

func TestRespond(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/profile", nil)

	someModel := models.Sessions{
		UserID:     "1",
		Hash:       "5",
		Expiration: 100,
	}

	handler := ProfilesServer{
		router: mux.NewRouter(),
		logger: logrus.New(),
		config: configs.NewConfig(),
	}

	handler.respond(w, r, 200, someModel)
	expected := http.StatusOK
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg, _ := json.Marshal(someModel)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestError(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/v1/profile", nil)

	someError := fmt.Errorf("some error")

	handler := ProfilesServer{
		router: mux.NewRouter(),
		logger: logrus.New(),
		config: configs.NewConfig(),
	}

	handler.error(w, r, 500, someError)
	expected := http.StatusInternalServerError
	if w.Code != expected {
		t.Errorf("expected: %v\n got: %v", expected, w.Code)
	}

	expectedMsg := fmt.Sprintf(`{"error":"%s"}`, someError)
	if !reflect.DeepEqual(string(expectedMsg), w.Body.String()) {
		t.Errorf("expected: %v\n got: %v", string(expectedMsg), w.Body.String())
	}
}

func TestCheckDBerrErrConstraintViolationEmail(t *testing.T) {
	errConstraintViolationEmail := profileModels.ErrConstraintViolationEmail

	reciveStr, reciveCode := checkDBerr(errConstraintViolationEmail)

	expectedStr := "Пользователь с таким email уже существует."
	expectedCode := http.StatusUnprocessableEntity

	if expectedStr != reciveStr {
		t.Errorf("expected: %v\n got: %v", expectedStr, reciveStr)
	}

	if expectedCode != reciveCode {
		t.Errorf("expected: %v\n got: %v", expectedCode, reciveCode)
	}
}

func TestCheckDBerrErrConstraintViolationNickname(t *testing.T) {
	errConstraintViolationNickname := profileModels.ErrConstraintViolationNickname

	reciveStr, reciveCode := checkDBerr(errConstraintViolationNickname)

	expectedStr := "Пользователь с таким nickname уже существует."
	expectedCode := http.StatusUnprocessableEntity

	if expectedStr != reciveStr {
		t.Errorf("expected: %v\n got: %v", expectedStr, reciveStr)
	}

	if expectedCode != reciveCode {
		t.Errorf("expected: %v\n got: %v", expectedCode, reciveCode)
	}
}

func TestCheckDBerrDefault(t *testing.T) {
	errDefault := fmt.Errorf("some strange error")

	reciveStr, reciveCode := checkDBerr(errDefault)

	expectedStr := "Неопознаная ошибка на севере, ухх..."
	expectedCode := http.StatusInternalServerError

	if expectedStr != reciveStr {
		t.Errorf("expected: %v\n got: %v", expectedStr, reciveStr)
	}

	if expectedCode != reciveCode {
		t.Errorf("expected: %v\n got: %v", expectedCode, reciveCode)
	}
}

func TestCheckDBerrJSONerr(t *testing.T) {
	JSONerr := fmt.Errorf(`{"error":"%s"}`, fmt.Sprint("some strange error"))

	reciveStr, reciveCode := checkDBerr(JSONerr)

	expectedStr := JSONerr.Error()
	expectedCode := http.StatusBadRequest

	if expectedStr != reciveStr {
		t.Errorf("expected: %v\n got: %v", expectedStr, reciveStr)
	}

	if expectedCode != reciveCode {
		t.Errorf("expected: %v\n got: %v", expectedCode, reciveCode)
	}
}
