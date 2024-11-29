package main

import (
	"fmt"
	"math"
)

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	if k1 == k2 {
		return math.NaN(), math.NaN()
	}

	x := (b2 - b1) / (k1 - k2)
	y := k1*x + b1

	return x, y
}

func main() {
	//var k1, b1, k2, b2 float64

	//fmt.Scanln(k1, b1, k2, b2)

	fmt.Println(FindIntersection(5, 20, -3, 10))
}
