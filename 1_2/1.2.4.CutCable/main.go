package main

import "fmt"

func CutCable(price []int, len int) int {
	dp := make([]int, len+1)

	// Find maximum value for all
	// rod of length i.
	for i := 1; i <= len; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = max(dp[i], price[j]+dp[i-j])
		}
	}

	fmt.Println(dp)
	return dp[len]
}

func main() {
	type Case struct {
		prices []int
		length int
		cost   int
	}

	Cases := []Case{
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

	for _, c := range Cases {
		cost := CutCable(c.prices, c.length)
		if cost != c.cost {
			fmt.Printf("Expected cost: %d, got cost: %d\n", c.cost, cost)
		} else {
			fmt.Printf("got cost: %d\n", cost)
		}
	}
}
