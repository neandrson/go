package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	var d float64

	fmt.Scanln(&a, &b, &c)

	d = b*b - 4.0*a*c
	//fmt.Println(d)
	if d > 0 {
		d1 := (-b + math.Sqrt(d)) / (2 * a)
		d2 := (-b - math.Sqrt(d)) / (2 * a)

		if d1 < d2 {
			fmt.Printf("%f %f", d1, d2)
		} else {
			fmt.Printf("%f %f", d2, d1)
		}
	} else if d == 0 {
		d1 := (-b + math.Sqrt(d)) / (2 * a)
		fmt.Printf("%f", d1)
	} else {
		d1 := 0
		d2 := 0
		fmt.Println(d1, d2)
	}
}

func main() {
	SqRoots()
}
