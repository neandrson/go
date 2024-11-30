package main

import "fmt"

func factorial(n int) int {
	// крайний случай
	if n == 0 {
		return 1
	}

	// вызов функции с текущим значением, умноженным на результат (n - 1)
	return n * factorial(n-1)
}

func main() {
	result := factorial(5)
	fmt.Println(result)
}
