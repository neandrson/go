package main

import (
	"os"
	"testing"
)

func TestLoadState(t *testing.T) {
	type test struct {
		cells     [][]bool
		fileName  string
		content   string
		wantError bool
	}

	tests := []test{
		{
			cells: [][]bool{
				{true, true},
				{false, false},
			},
			fileName: "file1",
			content:  "11\n00",
		},

		{
			cells: [][]bool{
				{true, true, false},
				{false, false, true},
			},
			fileName: "file2",
			content:  "110\n001",
		},
		{
			cells:     [][]bool{},
			fileName:  "file3",
			content:   "11\n0",
			wantError: true,
		},
	}

	for _, tt := range tests {
		if err := os.WriteFile(tt.fileName, []byte(tt.content), 0644); err != nil {
			t.Fatalf("TestLoadState failed on os.WriteFile with an error: %s", err)
		}
		defer os.Remove(tt.fileName)

		w := World{}
		if err := w.LoadState(tt.fileName); err != nil {
			if tt.wantError {
				continue
			}
			t.Fatalf("TestLoadState failed on file %q with an error: %s", tt.fileName, err)
		}

		/*if tt.wantError {
			t.Fatalf("TestLoadState expected %t, got %t", tt.wantError, t.Error)
		}*/
	}
}
