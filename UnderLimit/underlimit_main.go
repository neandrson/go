package main

import (
	"errors"
	"fmt"
	"testing"
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

func TestUnderLimit(t *testing.T) {
	type test struct {
		nums      []int
		n         int
		limit     int
		expected  []int
		wantError bool
	}

	tests := []test{
		{
			nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
			n:         5,
			limit:     3,
			expected:  []int{2, 0, 2},
			wantError: false,
		},
		{
			nums:      nil,
			wantError: true,
		},
		{
			nums:      []int{},
			n:         5,
			limit:     3,
			expected:  []int{},
			wantError: false,
		},
		{
			nums:      []int{3, 5, 6},
			n:         5,
			limit:     10,
			expected:  []int{3, 5, 6},
			wantError: false,
		},
		{
			nums:      []int{-13, 0, 6},
			n:         1,
			limit:     -5,
			expected:  []int{-13},
			wantError: false,
		},
		{
			nums:      []int{},
			n:         -1,
			limit:     5,
			wantError: true,
		},
	}

	for _, tc := range tests {
		result, err := UnderLimit(tc.nums, tc.limit, tc.n)
		if tc.wantError {
			if err == nil {
				t.Fatalf("expec... File is too long to be displayed fully")
			}
		}
	}
}

func main() {
	var b, c int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Scan("Введите числа общего и возращаемого массива ", &b, &c)

	d, err := UnderLimit(a, b, c)
	/*if err != nil {
		fmt.Println("Превышен объём")
	}*/
	fmt.Println(d, err)
}
