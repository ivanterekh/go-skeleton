package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestHello(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Logf("warning: could not init loger: %v", err)
	}
	router := setupRouter(logger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	code := w.Code
	expectedCode := 200
	if code != expectedCode {
		t.Errorf("unexpected status code: %v, want: %v", code, expectedCode)
	}

	response := w.Body.String()
	expectedResponse := "Hello, world!"
	if response != expectedResponse {
		t.Errorf("unexpected response: %v, want: %v", response, expectedResponse)
	}
}
