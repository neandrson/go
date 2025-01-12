package main

import (
	"testing"
)

func TestMaxExpressionValue(t *testing.T) {
	type testCase struct {
		nums []int
		val  int
	}

	testCases := []testCase{
		{
			nums: []int{3, 9, 10, 1, 30, 40},
			val:  46,
		},
		{
			nums: []int{4, 5, 10, 50, 25},
			val:  41,
		},
	}

	for _, tc := range testCases {
		val := MaxExpressionValue(tc.nums)
		if val != tc.val {
			t.Fatalf("Expected max value: %d, got: %d", tc.val, val)
		}
	}
}
