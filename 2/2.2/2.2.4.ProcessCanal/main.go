package main

import (
	"fmt"
	"slices"
)

func Process(nums []int) chan int {
	ch := make(chan int, 10)

	for _, num := range nums {
		ch <- num
	}
	return ch
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	ch := make(chan int, 10)

	ch = Process(nums)

	vals := []int{}

	for i := 0; i < 10; i++ {
		val := <-ch
		vals = append(vals, val)
	}

	slices.Sort(vals)

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !slices.Equal(vals, expected) {
		fmt.Printf("ch1 values expected: %v, got: %v\n", expected, vals)
	} else {
		fmt.Printf("ch1 values got: %v\n", vals)
	}
}
