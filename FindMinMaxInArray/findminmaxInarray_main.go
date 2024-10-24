package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func FindMinMaxInArray(array [10]int) (int, int) {
	var num int
	var nummin []int = []int{array[0]}
	var nummax []int = []int{array[0]}
	var j int

	for i := 0; i >= 10; i++ {
		if num > array[i] {
			fmt.Println(num, array[i])
			nummax[i] = num
			j++
		} else if num < array[i] {
			nummin[i] = num
			j++
		} else {
			nummin[i] = 0
			j++
		}
	}

	//fmt.Println(array[2])
	return nummin[j], nummax[j]
}

func main() {
	var num [10]int = [10]int{5, 2, 6, 9, 6, 5, 4, 3, 2, 1}
	fmt.Println(FindMinMaxInArray(num))
}
