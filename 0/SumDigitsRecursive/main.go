package main

import (
	"fmt"
)

func SumDigitsRecursive(n int) int {
	if n >= 0 && n < 10 {
		return n
	}
	if n < 0 {
		return SumDigitsRecursive(-n)
	}
	return SumDigitsRecursive(n%10) + SumDigitsRecursive(n/10)
}

func main() {
	fmt.Println(SumDigitsRecursive(123))
}
