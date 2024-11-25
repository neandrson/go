package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorizationMiddleware(t *testing.T) {
	req, err := http.NewRequest("GET", "/answer/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Добавляем заголовок Authorization с корректными данными
	req.Header.Set("Authorization", "Basic userid:password")

	rr := httptest.NewRecorder()

	handler := Authorization(answerHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expected := "The answer is 42"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %s, got %s", expected, rr.Body.String())
	}
}

func TestAuthorizationMiddlewareUnauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/answer/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Не добавляем заголовок Authorization

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Authorization(answerHandler))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	expected := "Unauthorized"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %s, got %s", expected, rr.Body.String())
	}
}
