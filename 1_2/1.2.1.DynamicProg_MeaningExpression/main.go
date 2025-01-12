package main

import "fmt"

func MaxExpressionValue(nums []int) int {
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i]) // Функция max — возвращает максимальное
	}

	second := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		second[i] = max(second[i+2], first[i+2]+nums[i+1]-nums[i])
	}
	// Здесь будет максимум
	return second[0]
}

func main() {
	nums := []int{3, 9, 10, 1, 30, 40}
	a := MaxExpressionValue(nums)
	fmt.Println(a)
}
