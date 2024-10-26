package main

import (
	"errors"
	"fmt"
)

func Clean(nums []int, x int) []int {
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
	/*var b, c int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Scan("Введите числа общего и возращаемого массива ", &b, &c)

	*/
	/*if err != nil {
		fmt.Println("Превышен объём")
	}*/
	d, err := TestUnderLimit()
	fmt.Println(d, err)
}
