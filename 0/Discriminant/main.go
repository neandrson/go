package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := -3.0
	c := 2.0

	discriminant := math.Pow(b, 2) - 4.0*a*c // math.Pow возводит b в степень 2
	fmt.Println(discriminant)                // Вывод: 1
}
