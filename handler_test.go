package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	router := setupRouter()

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
