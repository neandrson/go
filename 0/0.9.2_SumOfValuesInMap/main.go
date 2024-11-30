package main

import "fmt"

func SumOfValuesInMap(m map[int]int) int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}

func main() {
	m := map[int]int{
		10: 38,
		3:  19,
	}
	fmt.Println(SumOfValuesInMap(m))
}
