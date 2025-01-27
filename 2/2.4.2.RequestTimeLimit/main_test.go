package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}
func TestFetchAPI(t *testing.T) {

	go func() {
		http.HandleFunc("/hello", helloHandler)

		http.HandleFunc("/long", longHanlder)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			t.Errorf("error when starting a server")
		}
	}()

	time.Sleep(500)

	testCases := []struct {
		url     string
		timeout time.Duration
		want    APIResponse
		wantErr error
		name    string
	}{
		{
			name:    "ok",
			url:     "http://localhost:8080/hello",
			timeout: 2 * time.Second,
			want: APIResponse{
				Data:       `Hello, World!`,
				StatusCode: http.StatusOK,
			},
		},
		{
			name:    "timeout",
			url:     "http://localhost:8080/long",
			timeout: 10 * time.Millisecond,
			want: APIResponse{
				Data:    ``,
				StatusCode: http.StatusCode,
			},
		}
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			for _, element := range tc.elements {
				concurrentQueue.Enqueue(element)
			}

			for _, expected := range tc.elements {
				actual := concurrentQueue.Dequeue()
				if actual != expected {
					t.Errorf("Unexpected element. Got: %v, Expected: %v", actual, expected)
				}
			}

			var wg sync.WaitGroup

			for i := 1; i <= 10000; i++ {
				wg.Add(1)
				go func(item int) {
					concurrentQueue.Enqueue(item)
					wg.Done()
				}(i)
			}
			wg.Wait()

			if len(concurrentQueue.queue) != 10000 {
				t.Errorf("Unexpected len. Got: %v, Expected: %v", len(concurrentQueue.queue), tc.elements)
			}
		})
	}
}