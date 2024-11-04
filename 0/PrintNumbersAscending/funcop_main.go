package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {
	for i := 1; i < n; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(n)
}

func main() {
	var a, b float64
	var n int

	fmt.Scanln(&a, &b, &n)

	Add(a, b)
	Multiply(a, b)
	PrintNumbersAscending(n)
}
