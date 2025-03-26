package main

import "testing"

func TestSum(t *testing.T) {
	intSum := Sum([]int{1, 2, 3})
	if intSum != 6 {
		t.Errorf("Expected Sum([]int{1, 2, 3}) to be 6, but got %d", intSum)
	}

	floatSum := Sum([]float64{1.5, 2.5, 3.0})
	if floatSum != 7.0 {
		t.Errorf("Expected Sum([]float64{1.5, 2.5, 3.0}) to be 7.0, but got %f", floatSum)
	}
}
