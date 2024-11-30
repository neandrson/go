package main

import (
	"fmt"
)

func main() {
	var i int

	fmt.Scan(&i)
	if i < 0 && (i%2) == 0 {
		//fmt.Println(i % 2)
		fmt.Println("Число отрицательное и четное")
	} else if i < 0 && (i%2) <= -1 {
		//fmt.Println(i % 2)
		fmt.Println("Число отрицательное и нечетное")
	} else if i > 0 && (i%2) == 0 {
		//fmt.Println(i % 2)
		fmt.Println("Число положительное и четное")
	} else if i > 0 && (i%2) <= 1 {
		//fmt.Println(i % 2)
		fmt.Println("Число положительное и нечетное")
	} else {
		fmt.Println("Число равно нулю")
	}
}
