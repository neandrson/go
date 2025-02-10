package main

import (
	"fmt"
)

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	dp := 0
	if x == 0 {
		return 0
	}
	if x-s < 0 {
		dp = min(cs, min(cm, cl))
		return dp
	}

	if x-m < 0 {
		dp = min(cs+MinPizzaCost(s, m, l, cs, cm, cl, x-s), min(cm, cl))
		return dp
	}
	if x-l < 0 {
		dp = min(cl, min(cm+MinPizzaCost(s, m, l, cs, cm, cl, x-m), cs+MinPizzaCost(s, m, l, cs, cm, cl, x-s)))
		return dp
	}
	dp = min(cl+MinPizzaCost(s, m, l, cs, cm, cl, x-l), min(cm+MinPizzaCost(s, m, l, cs, cm, cl, x-m), cs+MinPizzaCost(s, m, l, cs, cm, cl, x-s)))
	return dp
}

func main() {
	type Case struct {
		s, m, l, cs, cm, cl, x int
		cost                   int
	}
	Cases := []Case{
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

	for _, c := range Cases {
		cost := MinPizzaCost(c.s, c.m, c.l, c.cs, c.cm, c.cl, c.x)
		if cost != c.cost {
			fmt.Printf("Expected cost: %d, got cost: %d\n", c.cost, cost)
		} else {
			fmt.Printf("got cost: %d\n", cost)
		}
	}
}
