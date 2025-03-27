package main

import (
	"testing"
)

func TestReadCSV(t *testing.T) {
	tests := []struct {
		name           string
		file           string
		expectedOutput [][]string
		expectedError  bool
	}{
		{
			name: "Empty CSV file",
			file: "empty.csv",
			// Assuming the file is empty, so no output is expected
			expectedOutput: nil,
			expectedError:  false,
		},
		{
			name: "Valid CSV file",
			file: "data.csv",
			// A sample expected output for testing purposes
			expectedOutput: [][]string{
				{"header1", "header2", "header3"},
				{"value1", "value2", "value3"},
			},
			expectedError: false,
		},
		{
			name: "No file",
			file: "data.csv",
			// A sample expected output for testing purposes
			expectedOutput: [][]string{},
			expectedError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if !test.expectedError {
				if err := creatCSV(test.file, test.expectedOutput); err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

			}
		})
	}
}
