package main

import "fmt"

func Contains(numbers []int, target int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Boolean: ", Contains(numbers, 5))
}
