package main

import (
	"testing"
)

func TestCut(t *testing.T) {
	type testCase struct {
		prices []int
		length int
		cost   int
	}

	testCases := []testCase{
		{
			prices: []int{0, 1, 5, 8, 9, 10, 17, 17, 20},
			length: 8,
			cost:   22,
		},
		{
			prices: []int{0, 3, 5, 6, 7, 10, 12},
			length: 6,
			cost:   18,
		},
	}

	for _, tc := range testCases {
		cost := CutCable(tc.prices, tc.length)
		if cost != tc.cost {
			t.Fatalf("Expected cost: %d, got cost: %d", tc.cost, cost)
		}
	}
}
