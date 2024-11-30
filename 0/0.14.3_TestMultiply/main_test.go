package main

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	if result != 6 {
		t.Errorf("Multiply(2, 3) = %d; want 6", result)
	}
	result = Multiply(5, 0)
	if result != 0 {
		t.Errorf("Multiply(5, 0) = %d; want 0", result)
	}

	result = Multiply(-4, 7)
	if result != -28 {
		t.Errorf("Multiply(-4, 7) = %d; want -28", result)
	}
}
