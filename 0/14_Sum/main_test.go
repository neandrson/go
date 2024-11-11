package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(5, 10)
	if result != 15 {
		t.Errorf("Sum function did not return the expected result. Got: %d, Want: %d", result, 15)
	}
}
