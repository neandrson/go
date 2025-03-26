package main

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	filtered := Filter([]int{1, 2, 3, 4, 5, 6}, even)
	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("Expected Filter([]int{1, 2, 3, 4, 5, 6}, even) to be %v, but got %v", expected, filtered)
	}
}
