package main

import (
	"net/http"
	"testing"

	"net/http/httptest"
)

func TestRequestHandlerBadRequestCase(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/greet", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("wrong status code")
	}
}
