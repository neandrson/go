package main

import (
	"testing"
)

func TestSortByFreq(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{
			str:      "abbbzzzat",
			expected: "taabbbzzz",
		},
		{
			str:      "adaabbcdd",
			expected: "cbbaaaddd",
		},
		{
			str:      "zggoooaaarrygyzv",
			expected: "vrryyzzaaagggooo",
		},
	}

	for _, tc := range tests {
		sorted := SortByFreq(tc.str)
		if sorted != tc.expected {
			t.Errorf("SortByFreq failed. Expected: %v, Got: %v", tc.expected, sorted)
		}

	}
}
