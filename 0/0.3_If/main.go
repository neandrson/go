package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scanln(&i)
	if i > 0 {
		fmt.Println("Число положительное")
	} else if i < 0 {
		fmt.Println("Число отрицательное")
	} else if i == 0 {
		fmt.Println("Введен ноль")
	}
}
