package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a / b), nil
}

func DivideIntegers(a, b int) (float64, error) {
	result, err := divide(a, b)
	if err != nil {
		//fmt.Println("Ошибка:", err)
		return 0, err
	}
	//fmt.Println("Результат:", result)
	return result, nil

}

func main() {
	fmt.Println(DivideIntegers(10, 2))
}
