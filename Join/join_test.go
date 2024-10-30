package main

import (
	"slices"
	"testing"
)

func TestJoin(t *testing.T) {
	type test struct {
		nums1    []int
		nums2    []int
		expected []int
	}

	tests := []test{
		{
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums1:    []int{},
			nums2:    []int{4, 5, 6, 7},
			expected: []int{4, 5, 6, 7},
		},
		{
			nums1:    []int{4, 5, 6, 7},
			nums2:    []int{},
			expected: []int{4, 5, 6, 7},
		},
		{
			nums1:    []int{},
			nums2:    []int{},
			expected: []int{},
		},
		{
			nums1:    []int{4, 6, 149, 5, 2, 1},
			nums2:    []int{99},
			expected: []int{4, 6, 149, 5, 2, 1, 99},
		},
	}

	for _, tc := range tests {
		res := Join(tc.nums1, tc.nums2)

		if cap(res) != cap(tc.nums1)+cap(tc.nums2) {
			t.Fatalf("expected cap: %v, got cap: %v",
				cap(tc.nums1)+cap(tc.nums2), cap(res))
		}

		if !slices.Equal(res, tc.expected) {
			t.Fatalf("expected: %v, got: %v", tc.expected, res)
		}

	}
}
