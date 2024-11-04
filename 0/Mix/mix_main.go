package main

import "fmt"

func Mix(nums []int) []int {
	var nums_half = len(nums) / 2

	var result []int
	for i := 0; i < nums_half; i++ {
		result = append(result, nums[i], nums[nums_half+i])
	}

	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 11, 22, 33, 44, 55, 66, 77}

	fmt.Println(a)
	fmt.Println(Mix(a))
}
