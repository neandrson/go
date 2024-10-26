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

/*func TestUnderLimit(t *testing.T) {
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
				t.Fatalf("n cannot be negative")
			}
		}
	}
}*/

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
