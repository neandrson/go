package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func FiveSteps(array [5]int) [5]int {
	var numb [5]int
	var j int

	for i := 4; i >= 0; i-- {
		numb[j] = array[i]
		j++
	}
	return numb
}

func main() {
	var num [5]int = [5]int{5, 4, 3, 2, 1}
	fmt.Println(FiveSteps(num))
}
