package main

import (
	"fmt"
	"reflect"
)

type MyFilters interface {
	int | float64
}

func Filter[T MyFilters](inputs []T, fn func(T) bool) []T {
	var outputs []T
	for _, input := range inputs {
		if fn(input) {
			outputs = append(outputs, input)
		}
	}
	return outputs
}

func main() {
	even := func(x int) bool { return x%2 == 0 }
	filtered := Filter([]int{1, 2, 3, 4, 5, 6}, even)
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(filtered, expected) {
		fmt.Printf("Expected Filter([]int{1, 2, 3, 4, 5, 6}, even) to be %v, but got %v", expected, filtered)
	} else {
		fmt.Printf("Filter([]int{1, 2, 3, 4, 5, 6}, even) got %v", filtered)
	}
}
