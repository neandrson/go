package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

func ThirdElementInArray(array [6]int) int {
	//var numb [6]int = [6]array
	//fmt.Println(array[2])
	return array[2]
}

func main() {
	var num [6]int = [6]int{6, 5, 4, 3, 2, 1}
	fmt.Println(ThirdElementInArray(num))
}
