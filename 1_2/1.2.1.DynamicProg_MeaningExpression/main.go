package main

import "fmt"

func MaxExpressionValue(nums []int) int {
	first := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		first[i] = max(first[i+1], nums[i]) // Функция max — возвращает максимальное
	}

	second := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		second[i] = max(second[i+1], first[i+1]-nums[i])
	}

	thry := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		thry[i] = max(thry[i+1], second[i+1]+nums[i])
	}

	fhor := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		fhor[i] = max(fhor[i+1], thry[i+1]-nums[i])
	}
	// Здесь будет максимум
	return fhor[0]
}

func main() {
	nums := []int{4, 5, 10, 50, 25}
	a := MaxExpressionValue(nums)
	fmt.Println(a)
}
