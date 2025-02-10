package main

import (
	"testing"
)

func TestMinPizzaCost(t *testing.T) {
	type testCase struct {
		s, m, l, cs, cm, cl, x int
		cost                   int
	}
	testCases := []testCase{
		{
			s:    4,
			m:    6,
			l:    12,
			cs:   40,
			cm:   60,
			cl:   100,
			x:    17,
			cost: 160,
		},
		{
			s:    3,
			m:    6,
			l:    9,
			cs:   50,
			cm:   150,
			cl:   300,
			x:    16,
			cost: 300,
		},
	}

	for _, tc := range testCases {
		cost := MinPizzaCost(tc.s, tc.m, tc.l, tc.cs, tc.cm, tc.cl, tc.x)
		if cost != tc.cost {
			t.Fatalf("Expected cost: %d, got cost: %d", tc.cost, cost)
		}
	}
}
