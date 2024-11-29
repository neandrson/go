package main

import (
	"fmt"
	"strconv"
)

func CalculateSeriesSum(n int) float64 {
	var j float64
	for i := 1; i <= n; i++ {
		//fmt.Println(i, n)
		j += 1 / float64(i)
	}
	k := fmt.Sprintf("%.5f", j)
	l, _ := strconv.ParseFloat(k, 64)
	return l
}

func main() {
	//var i float64
	//i =
	fmt.Println(CalculateSeriesSum(6))
}
