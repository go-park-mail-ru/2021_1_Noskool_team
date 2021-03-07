package middleware

import (
	"2021_1_Noskool_team/configs"
	mock_client "2021_1_Noskool_team/internal/microservices/auth/delivery/grpc/client/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})
	handlerToTest := LoggingMiddleware(nextHandler)

	req := httptest.NewRequest("GET", "/api/v1/", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, req)
}

func TestCORSMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.True(t, true)
	})

	corsMid := NewCORSMiddleware(&configs.Config{
		FrontendUrl: "some url",
	})
	handlerToTest := corsMid.CORS(nextHandler)

	req := httptest.NewRequest("GET", "/api/v1/", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, req)
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Methods"),
		"POST, GET, OPTIONS, PUT, DELETE")
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Origin"),
		corsMid.config.FrontendUrl)
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Credentials"),
		"true")
}

func TestPanicMiddleware(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})
	handlerToTest := PanicMiddleware(nextHandler)
	req := httptest.NewRequest("GET", "/api/v1/", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, req)
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
