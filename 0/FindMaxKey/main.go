package main

import "fmt"

func FindMaxKey(m map[int]int) int {
	var maxNumber int
	for n := range m {
		maxNumber = n
		break
	}
	for n := range m {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}

func main() {
	m := map[int]int{
		-7: 1,
		-1: 38,
	}
	fmt.Println(FindMaxKey(m))
}
