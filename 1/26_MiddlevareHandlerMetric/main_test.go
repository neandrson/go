package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibonacciHandler(t *testing.T) {
	// Clear the request count before testing
	requestCount = 0

	tt := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{"First Fibonacci Number", http.StatusOK, "0"},
		{"Second Fibonacci Number", http.StatusOK, "1"},
		{"Third Fibonacci Number", http.StatusOK, "1"},
		{"Fourth Fibonacci Number", http.StatusOK, "2"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(FibonacciHandler)

			handler.ServeHTTP(rr, req)

			if tc.expectedStatus != rr.Code {
				t.Errorf("failed error code")
			}

			if tc.expectedBody != rr.Body.String() {
				t.Errorf("failed body")
			}
		})
	}
}

func TestMetricsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MetricsHandler)

	handler.ServeHTTP(rr, req)

	/*if t.expectedStatus != rr.Code {
		t.Errorf("failed error code")
	}

	if t.expectedBody != rr.Body.String() {
		t.Errorf("failed body")
	}*/
}
