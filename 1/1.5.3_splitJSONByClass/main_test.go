package main

import (
	"bytes"
	"testing"
)

func TestSplitJSONByClass(t *testing.T) {
	inputJSON := []byte(`[
        {
            "name": "Oleg",
            "class": "9B"
        },
        {
            "name": "Ivan",
            "class": "9A"
        },
        {
            "name": "Maria",
            "class": "9B"
        },
        {
            "name": "John",
            "class": "9A"
        }
    ]`)

	expectedJSONMap := map[string][]byte{
		"9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
		"9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]`),
	}

	classJSONMap, err := splitJSONByClass(inputJSON)
	if err != nil {
		t.Fatalf("Error while splitting JSON by class: %v", err)
	}

	for class, expectedJSON := range expectedJSONMap {
		if !bytes.Equal(classJSONMap[class], expectedJSON) {
			t.Errorf("Expected JSON data for class %s to be %s, but got %s", class, expectedJSON, classJSONMap[class])
		}
	}
}
