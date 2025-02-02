package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var sleepTime time.Duration

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(sleepTime)
	fmt.Fprintf(w, "Hello, World!")
}
func TestReadSourceHandler(t *testing.T) {

	go func() {
		http.HandleFunc("/provideData", longHanlder)

		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			t.Errorf("error when starting a server")
		}
	}()
	// start on 8080
	go StartServer(time.Second)

	time.Sleep(100)
	tests := []struct {
		name         string
		timeout      time.Duration
		expectedCode int
		expectedBody string
	}{
		{
			name:         "ValidRequest",
			timeout:      time.Microsecond,
			expectedCode: http.StatusOK,
			expectedBody: "Hello, World!",
		},
		{
			name:         "TimeoutRequest",
			timeout:      time.Second * 5,
			expectedCode: http.StatusServiceUnavailable,
			expectedBody: "Request timeout",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			sleepTime = test.timeout
		})
	}
}
