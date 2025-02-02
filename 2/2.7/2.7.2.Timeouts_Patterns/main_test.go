package main

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestReadJSONWithTimeout(t *testing.T) {
	// create a temporary test JSON file
	testData := `{"message": "What is Yandex Lyceum?"}` // replace this with your test JSON data
	tempFile, err := ioutil.TempFile("", "test*.json")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // clean up the temporary file
	defer tempFile.Close()

	// write test JSON data to the temporary file
	_, err = tempFile.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	// create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// create a channel for the result
	resultChan := make(chan []byte)

	// call readJSON function in a Goroutine
	go readJSON(ctx, tempFile.Name(), resultChan)

	select {
	case data := <-resultChan:
		// reading JSON data successful, check the content
		
		default

	}
}