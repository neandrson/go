package main

import (
	"fmt"
	"slices"
)

func Clean(nums []int, x int) []int {
	return slices.DeleteFunc(nums, func(v int) bool {
		return v == x
	})
}

func main() {
	//var b, c int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 3
	/*fmt.Scan("Введите числа общего и возращаемого массива ", &b, &c)

	 */
	/*if err != nil {
		fmt.Println("Превышен объём")
	}*/
	b := Clean(a, n)
	fmt.Println(b)
}
