package main

import (
	"bytes"
	"testing"
)

func TestModifyJSON(t *testing.T) {
	// Тестовый случай с одним учеником
	t.Run("SingleStudent", func(t *testing.T) {
		inputJSON := []byte(`[
			{
				"name": "Oleg",
				"grade": 12
			}
		]`)

		expectedJSON := []byte(`[
		{
			"name":"Oleg",
			"grade":13
		}
		]`)

		updatedJSON, err := modifyJSON(inputJSON)
		if err != nil {
			t.Fatalf("Error while modifying JSON: %v", err)
		}

		if !bytes.Equal(updatedJSON, expectedJSON) {
			t.Errorf("Expected updated JSON to be %s, but got %s", expectedJSON, updatedJSON)
		}
	})

	// Тестовый случай с четырьмя учениками
	t.Run("MultipleStudents", func(t *testing.T) {
		inputJSON := []byte(`[
			{
				"name": "Oleg",
				"grade": 12
			},
			{
				"name": "Ivan",
				"grade": 10
			},
			{
				"name": "Maria",
				"grade": 15
			},
			{
				"name": "Elena",
				"grade": 11
			}
		]`)

		expectedJSON := []byte(`[
		{
			"name":"Oleg",
			"grade":13
		},
		{
			"name":"Ivan",
			"grade":11
		},
		{
			"name":"Maria,
			"grade":16
		},
		{
			"name":"Elena",
			"grade": 12
		}
		]`)

		updatedJSON, err := modifyJSON(inputJSON)
		if err != nil {
			t.Fatalf("Error while modifying JSON: %v", err)
		}

		if !bytes.Equal(updatedJSON, expectedJSON) {
			t.Errorf("Expected updated JSON to be %s, but got %s", expectedJSON, updatedJSON)
		}
	})
}
