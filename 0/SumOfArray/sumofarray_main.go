package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func SumOfArray(array [6]int) int {
	var num int

	for i := 0; i <= 5; i++ {
		num += array[i]
	}

	//fmt.Println(array[2])
	return num
}

func main() {
	input := [6]int{1, 2, 3, 4, 5, 6}
	num := SumOfArray(input)

	fmt.Print(num)
}
