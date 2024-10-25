package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func PrettyArrayOutput(array [9]string) {
	var nummin, nummax int

	for i := 0; i <= 8; i++ {
		fmt.Println(array[i])
		if array[i] == 0 {
			fmt.Println(array[i])
			nummin = 0
		} else if nummax < array[i] {
			fmt.Println(nummax)
			nummax = array[i]
		} else {
			nummin = array[i]
			fmt.Println(nummin)

		}
	}

	//fmt.Println(array[2])
	return nummax, nummin
}

func main() {
	input := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max, min := FindMinMaxInArray(input)

	fmt.Print(max, min)
}
