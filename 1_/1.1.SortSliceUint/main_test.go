package main

import (
	"slices"
	"testing"
)

func TestSortNums(t *testing.T) {
	tests := []struct {
		nums     []uint
		expected []uint
	}{
		{
			nums:     []uint{4, 1, 5, 0},
			expected: []uint{0, 1, 4, 5},
		},
		{
			nums:     []uint{490, 741, 88, 1, 10, 7, 234, 2234, 64, 12, 778, 21234, 4345, 45673, 23, 5, 78, 2, 1, 5},
			expected: []uint{1, 1, 2, 5, 5, 7, 10, 12, 23, 64, 78, 88, 234, 490, 741, 778, 2234, 4345, 21234, 45673},
		},
	}
	for _, tc := range tests {
		SortNums(tc.nums)
		if slices.Compare(tc.expected, tc.nums) != 0 {
			t.Errorf("TestSortNums failed. Expected: %v, Got: %v", tc.expected, tc.nums)
		}
	}

}
