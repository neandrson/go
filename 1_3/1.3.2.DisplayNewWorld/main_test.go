package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	height := 3
	width := 3

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	w := &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}

	representation := fmt.Sprint(w)

	if strings.Contains(representation, "false") {
		t.Fatalf("You need to implement String() method")
	}
}
