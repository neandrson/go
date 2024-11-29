package main

import "fmt"

var memo = map[int]int{}

func fib(n int) int {
	if n < 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

func main() {
	fmt.Println(fib(12345))
}
