package main

import (
	"fmt"
	"net/http"
)

func FetchURL(url string) string {
	//client := http.Client{Timeout: 5 * time.Second}

	//получение ответа по URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Failed to fetch")
	}
	defer resp.Body.Close()
	return fmt.Sprintf("Successfully fetched")
}

func main() {
	expected := "Failed to fetch"

	result := FetchURL("https://ya.ru")

	if result != expected {
		fmt.Printf("Expected '%s', but got '%s'", expected, result)
	}
}
