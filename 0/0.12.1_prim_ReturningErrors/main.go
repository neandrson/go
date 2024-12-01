package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

/*func divide(a, b float64) (float64, error) {
	if b == 0 {
		//return 0, fmt.Errorf("division by zero")
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}*/

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", result)
}
