package main

import "errors"

type World struct {
	Height int
	Width  int
	Cells  [][]bool
}

func NewWorld(height, width int) (*World, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("Height and width must be positive numbers")
	}
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}

	return &World{
		Height: height,
		Width:  width,
		Cells:  cells,
	}, nil
}

func main() {

}
