package main

import (
	"testing"
)

// creating a test function
func TestFetchURL(t *testing.T) {

	expected := "Failed to fetch"

	result := FetchURL("https://ya.ru")

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
