package middleware

import (
	"2021_1_Noskool_team/configs"
	mock_client "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client/mocks"
	"github.com/BurntSushi/toml"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	configPath = "configs/config.toml"
)

func TestCORSMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})

	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Error(err)
	}

	corsMid := NewCORSMiddleware(config)
	handlerToTest := corsMid.CORS(nextHandler)

	req := httptest.NewRequest("GET", "/api/v1/", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, req)
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Methods"),
		"POST, GET, OPTIONS, PUT, DELETE")
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Origin"),
		config.FrontendURL)
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Credentials"),
		"true")
}

func TestCheckSessionMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//result := models.Result{
	//	ID:     1,
	//	Status: "OK",
	//}
	client := mock_client.NewMockAuthCheckerClient(ctrl)
	//client.EXPECT().Check(context.Background(),
	//	result.ID).Times(1).Return(result, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/api/vi/profiles/", nil)
	//r = mux.SetURLVars(r, map[string]string{"session_id": ""})
	//r.Header.Add("session_id", "1")

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
	sessionsMidl := NewSessionMiddleware(client)
	handlerToTest := sessionsMidl.CheckSessionMiddleware(nextHandler)
	handlerToTest.ServeHTTP(w, r)
}

func TestContentTypeJson(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})

	req := httptest.NewRequest("GET", "/api/v1/", nil)
	w := httptest.NewRecorder()

	handlerToTest := ContentTypeJson(nextHandler)
	handlerToTest.ServeHTTP(w, req)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
}
