package main

import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	// Проверка на отрицательное значение n
	if n < 0 {
		return nil, errors.New("n cannot be negative")
	}
	// Инициализация результирующего слайса
	result := []int{}
	// Перебор элементов в слайсе
	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			if len(result) == n {
				break
			}
		}
	}
	return result, nil
}

func main() {
	nums := []int{5, 3, 10, 7, 2, 8}
	limit := 6
	n := 3

	result, err := UnderLimit(nums, limit, n)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
