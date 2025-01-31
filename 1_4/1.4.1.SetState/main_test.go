package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(time.Second * 2)
	os.Exit(m.Run())
}

func TestServer(t *testing.T) {
	sendRequest(t, 30)
	sendRequest(t, 50)
}

func sendRequest(t *testing.T, percentage int) {
	fillPercentage := struct {
		Fill int `json:"fill"`
	}{
		Fill: percentage,
	}
	reqBody, err := json.Marshal(fillPercentage)
	if err != nil {
		t.Fatal(err)
	}

	// Send the HTTP POST request
	resp, err := http.Post("http://localhost:8081/setstate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Request failed with status code %d", resp.StatusCode)
	}

	// Check the file content
	content, err := os.ReadFile("state.cfg")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("state.cfg")

	fileBody := string(content)

	if fileBody != fmt.Sprintf("%d%%", "percentage", fileBody, "") {
		t.Fatalf("%d%%", "percentage", fileBody, "")
	}
}
