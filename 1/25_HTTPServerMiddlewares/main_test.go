package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tt := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{"Valid Name", "Alice", http.StatusOK, "Hello, Alice!"},
		{"Empty Name", "", http.StatusOK, "Hello, stranger!"},
		{"Invalid Name", "John$Doe", http.StatusOK, "Hello, dirty hacker!"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/hello?name="+tc.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}
			//fmt.Println(req)

			rr := httptest.NewRecorder()
			handler := SetDefaultName(Sanitize(HelloHandler))
			//fmt.Println(handler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatus, rr.Code)
			}

			if rr.Body.String() != tc.expectedBody {
				t.Errorf("Expected response body \"%s\", but got \"%s\"", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
