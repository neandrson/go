package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	go main()
	time.Sleep(time.Second * 2)
	os.Exit(m.Run())
}
func TestServer(t *testing.T) {
	url := "http://localhost:8081/reset"
	method := "PUT"

	makeRequest(t, 10, method, url)

	makeRequest(t, 60, method, url)
}

func saveFillPercentage(fill int) error {
	percentage := fmt.Sprintf("%d%%", fill)
	err := os.WriteFile("state.cfg", []byte(percentage), 0644)
	if err != nil {
		return err
	}
	return nil
}

func makeRequest(t *testing.T, percentage int, method, url string) {
	payload := strings.NewReader("")

	saveFillPercentage(percentage)
	defer os.Remove("state.cfg")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		t.Fatal(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
		return
	}
	defer res.Body.Close()

	// Check the response status co
}
