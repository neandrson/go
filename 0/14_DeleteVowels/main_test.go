package main

import "testing"

func TestDeleteVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hll"},
		{"world", "wrld"},
		{"apple", "ppl"},
		{"banana", "bnn"},
		{"", ""},
		{"elephant", "lphnt"},
		{"orange", "rng"},
	}

	for _, test := range tests {
		result := DeleteVowels(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
		}
	}
}
