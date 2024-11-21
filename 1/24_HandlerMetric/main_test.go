package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestFib(t *testing.T) {
	fibSlice := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for i, fib := range fibSlice {
		req, err := http.NewRequest("GET", "/metrics", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		MetricsHandler(w, req)
		res := w.Result()
		defer res.Body.Close()

		body2, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("server works incorrect")
		}
		if "rpc_duration_milliseconds_count "+strconv.Itoa(i) != strings.TrimSpace(string(body2)) {
			t.Fatal("incorrect metric", string(body2), "expect", i)
		}
		// -------------
		req, err = http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w = httptest.NewRecorder()
		FibonacciHandler(w, req)
		res = w.Result()
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal("server works incorrect")
		}
		if strconv.Itoa(fib) != strings.TrimSpace(string(body)) {
			t.Fatal("incorrect metric", string(body), "expect", i)
		}
	}
}
