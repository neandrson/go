package main

import (
	"fmt"
)

func Multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	return a * b
}

func main() {
	fmt.Println(Multiply(5, 5))
}
