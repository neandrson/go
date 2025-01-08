package main

import (
	"slices"
	"testing"
)

func TestSortAndMerge(t *testing.T) {
	tests := []struct {
		left     []int
		right    []int
		expected []int
	}{
		{
			left:     []int{4, 1, 5, 0},
			right:    []int{-1, 4, 5, 10},
			expected: []int{-1, 0, 1, 4, 4, 5, 5, 10},
		},
		{
			left:     []int{490, 741, 88, 1, 10, 7, 234, 2234, 64, -12, 778, 21234, 4345, 45673, -23, 5, 78, 2, 1, 5},
			right:    []int{-1, 4, 5, 104, 1, 18733, 0},
			expected: []int{-23, -12, -1, 0, 1, 1, 1, 2, 4, 5, 5, 5, 7, 10, 64, 78, 88, 104, 234, 490, 741, 778, 2234, 4345, 18733, 21234, 45673},
		},
	}
	for _, tc := range tests {
		merged := SortAndMerge(tc.left, tc.right)
		if slices.Compare(tc.expected, merged) != 0 {
			t.Errorf("TestSortNums failed. Expected: %v, Got: %v", tc.expected, merged)
		}
	}
}
