package main

import "fmt"

func fib(n int) int {
	// крайний случай, так как при f(1) и так далее мы не сможем посчитать обе следующих итерации
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(6)) // Вывод: 8
}
