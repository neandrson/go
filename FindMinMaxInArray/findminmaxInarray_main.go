package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func FindMinMaxInArray(array [10]int) (int, int) {
	var nummin, nummax int
	nummax = array[0]
	nummin = array[0]
	for i := 0; i <= 9; i++ {
		if nummax < array[i] {
			fmt.Println(array[i])
			nummax = array[i]
		}
		if nummin > array[i] {
			fmt.Println(array[i])
			nummin = array[i]
		}
	}

	//fmt.Println(array[2])
	return nummax, nummin
}

func main() {
	input := [10]int{-10, -8, -9, -7, -5, -4, -6, -3, -1, -2}
	max, min := FindMinMaxInArray(input)

	fmt.Print(max, min)
}
