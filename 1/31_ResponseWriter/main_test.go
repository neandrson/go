package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"log"
)

func TestDefault(t *testing.T) {
	// Create a buffer to capture logs
	var buf bytes.Buffer

	// Set the logger's output to the buffer
	log.SetOutput(&buf)

	type testresponse struct {
		Name string `json:"name"`
	}

	input := []string{"john", "jack", "ivan"}

	for i := 0; i < len(input); i++ {
		req, err := http.NewRequest("GET", "/?name="+input[i], nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		HelloHandler(w, req)
		resp := w.Result()
		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("server works incorrect")
		}

		expectedResponse := testresponse{Name: input[i]}
		b, _ := json.Marshal(expectedResponse)

		logOutput := buf.String()
		res := strings.Split(logOutput, " ")
		expectedRes := strings.Trim(res[len(res)-1], "\n")
		if string(b) != expectedRes {
			t.Fatal(string(b), " not ", expectedRes)
		}
	}
}
