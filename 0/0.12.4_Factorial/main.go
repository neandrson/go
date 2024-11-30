package main

import (
	"errors"
	"fmt"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}
	return factorial, nil
}

func main() {
	fmt.Println(Factorial(5))
}
