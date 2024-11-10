package main

import (
	"errors"
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	return n * Factorial(n-1), nil
}

func main() {
	fmt.Println(Factorial(5))
}
