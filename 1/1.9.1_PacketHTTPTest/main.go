package main

import (
	"fmt"
	"io"

	"net/http"

	"net/http/httptest"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	req := httptest.NewRequest("GET", "http://google.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("body: ", string(body))
	fmt.Println("status code: ", resp.StatusCode)
}
