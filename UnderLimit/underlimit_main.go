package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	a := nums[:]
	//var err errors
	if len(a) < n; limit, err = int.Pars {
		return a[:limit], errors.ErrUnsupported
	}
	return a[:n], nil

}

func main() {
	var b, c int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Scan("Введите общую и количество возращаемого массива ", &b, &c)
	
	d := make([]int, b)
	if err != nil {
		fmt.Println("Превышен объём")
	}
	fmt.Println(d)
}
