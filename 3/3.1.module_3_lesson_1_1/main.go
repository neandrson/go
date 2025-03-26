package main

import "fmt"

type MyConstraint interface {
	int | float64
}

func Sum[T MyConstraint](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	intSum := Sum([]int{1, 2, 3})
	if intSum != 6 {
		fmt.Printf("Expected Sum([]int{1, 2, 3}) to be 6, but got %d\n", intSum)
	} else {
		fmt.Printf("Sum([]int{1, 2, 3}) got %d\n", intSum)
	}

	floatSum := Sum([]float64{1.5, 2.5, 3.0})
	if floatSum != 7.0 {
		fmt.Printf("Expected Sum([]float64{1.5, 2.5, 3.0}) to be 7.0, but got %f\n", floatSum)
	} else {
		fmt.Printf("Sum([]float64{1.5, 2.5, 3.0}) got %f\n", floatSum)
	}
}
