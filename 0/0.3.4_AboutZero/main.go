package main

import (
	"fmt"
)

func main() {
	var i int
	var j int

	fmt.Scan(&i)
	fmt.Scan(&j)
	if i > 0 && j > 0 {
		//fmt.Println(i % 2)
		fmt.Println("Оба числа положительные")
	} else if i < 0 && j < 0 {
		//fmt.Println(i % 2)
		fmt.Println("Оба числа отрицательные")
	} else if (i > 0 && j < 0) || (i < 0 && j > 0) {
		//fmt.Println(i % 2)
		fmt.Println("Одно число положительное, а другое отрицательное")
	} else if i == 0 || j == 0 {
		//fmt.Println(i % 2)
		fmt.Println("Одно из чисел равно нулю")
	}
}
