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

func hiHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Microsecond * 100)
	fmt.Fprintf(w, "Hi, World!")
}

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello, World!")
}
func TestFetchAPI(t *testing.T) {
	go func() {
		http.HandleFunc("/hello", helloHandler)

		http.HandleFunc("/long", longHanlder)

		http.HandleFunc("/hi", hiHandler)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			t.Errorf("error when starting a server")
		}
	}()

	time.Sleep(100)

	testCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type testStruct struct {
		ctx      context.Context
		urls     []string
		timeout  time.Duration
		expected []*APIResponse
		err      error
	}

	tests := []testStruct{
		{
			ctx:     testCtx,
			urls:    []string{"http://localhost:8080/hello","http://localhost:8080/long","http://localhost:8080/hi"},
			timeout: 2* time.Second,
			expected: []*APIResponse{
				URL        string // запрошенный URL
				Data       string // тело ответа
				StatusCode int    // код ответа
				Err        error
			}
		}
	}
	
}