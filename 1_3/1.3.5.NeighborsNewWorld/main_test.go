package main

import (
	"testing"
)

func TestNeighbors(t *testing.T) {

	type test struct {
		cells     [][]bool
		X         int
		Y         int
		neighbors int
	}

	tests := []test{
		{
			cells: [][]bool{
				{true, false, true},
				{false, false, false},
				{true, false, false},
			},
			X:         0,
			Y:         0,
			neighbors: 2,
		},
		{
			cells: [][]bool{
				{true, false, false},
				{false, false, true},
				{true, false, false},
			},
			X:         2,
			Y:         1,
			neighbors: 2,
		},
	}

	for _, tt := range tests {
		w := World{
			Height: len(tt.cells),
			Width:  len(tt.cells[0]),
			Cells:  tt.cells,
		}
		n := w.Neighbors(tt.X, tt.Y)

		if n != tt.neighbors {
			t.Fatalf("expected: %v, got: %v", tt.neighbors, n)
		}
	}
}
