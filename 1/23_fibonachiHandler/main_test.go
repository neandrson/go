package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	go StartServer(time.Second)
	// time.Sleep(5 * time.Second)
	fibSlice := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i, fib := range fibSlice {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		FibonacciHandler(w, req)
		res := w.Result()
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("server works incorrect", err)
		}
		if strconv.Itoa(fib) != strings.TrimSpace(string(body)) {
			t.Fatal("incorrect answer", i, string(body), "expect", fib)
		}
	}
}
