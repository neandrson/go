package main

import (
	"fmt"
)

type Chest struct {
	val        []int
	wt         []int
	itemsToPut []int
	cap        int
	cost       int
}

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	vales := Chest{val: []int{s, m, l}, wt: []int{cs, cm, cl}}
	// рассчет стоимости
	n := len(vales.val)

	// create dp data structure
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= n; j++ {
			if vales.wt[i] > j {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = min(dp[i][j-vales.wt[i]]+vales.val[i], dp[i][j])
			}
		}
	}
	fmt.Println(dp, n)
	return dp[n][m]
}

func main() {
	s := 314
	m := 706
	l := 1256
	cs := 230
	cm := 510
	cl := 925
	x := 500
	a := MinPizzaCost(s, m, l, cs, cm, cl, x)
	fmt.Println(a)
}
