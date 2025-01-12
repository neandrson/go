package main

import (
	"testing"

	"slices"
)

func TestKnapsackEasy(t *testing.T) {

	type testCase struct {
		itemsCost   []int
		itmesWeight []int
		itemsToPut  []int
		cap         int
		cost        int
	}

	testCases := []testCase{
		{
			itemsCost:   []int{100, 400, 300, 500},
			itmesWeight: []int{5, 4, 6, 3},
			itemsToPut:  []int{1, 3},
			cap:         10,
			cost:        900,
		},
		{
			itemsCost:   []int{5, 3, 4},
			itmesWeight: []int{3, 2, 1},
			itemsToPut:  []int{0, 2},
			cap:         5,
			cost:        9,
		},
	}

	for _, tc := range testCases {

		chest := &Chest{
			val: tc.itemsCost,
			wt:  tc.itmesWeight,
		}
		cost, items := Knapsack(chest, tc.cap)

		if cost != tc.cost {
			t.Fatalf("Expected cost: %d, got cost: %d", tc.cost, cost)
		}

		slices.Sort(items)
		slices.Sort(tc.itemsToPut)
		if !slices.Equal(items, tc.itemsToPut) {
			t.Fatalf("Expected items: %v, got items: %v", tc.itemsToPut, items)
		}
	}
}
