package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tt := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{"Valid Name", "Alice", http.StatusOK, `{"status":"ok","result":{"greetings":"hello","name":"Alice"}}`},
		{"Empty Name", "", http.StatusOK, `{"status":"ok","result":{"greetings":"hello","name":"stranger"}}`},
		{"Invalid Name", "John$Doe", http.StatusOK, `{"status":"error","result":{}}`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/hello?name="+tc.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := RPC(SetDefaultName(Sanitize(HelloHandler)))

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, rr.Code)
			}

			if strings.TrimSpace(rr.Body.String()) != tc.expectedBody {
				t.Errorf("Expected status code %s, but got %s", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
