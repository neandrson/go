package main

import (
	"fmt"
	"math"
)

/*func main() {
	a := 1.0
	b := -3.0
	c := 2.0

	discriminant := math.Pow(b, 2) - 4.0*a*c
	fmt.Println(discriminant) // Вывод: 1

	a = 2.0
	b = 4.0
	c = 2.0

	discriminant = math.Pow(b, 2) - 4.0*a*c
	fmt.Println(discriminant) // Вывод: 0
}*/

/*func findDiscriminant(a, b, c float64) {
	// тело функции
	discriminant := math.Pow(b, 2) - 4.0*a*c
	fmt.Println(discriminant)
}

func main() {
	findDiscriminant(1, -3.0, 2) // a = 1  b = -3.0  c = 2
	// Вывод: 1

	findDiscriminant(1, 4.0, 2) // a = 1  b = 4.0   c = 2
	// Вывод: 8
}*/

func findDiscriminant(a, b, c float64) float64 {
	discriminant := math.Pow(b, 2) - 4.0*a*c

	// возвращаем значения вычисления из findDiscriminant
	return discriminant
}

func main() {
	// в данном случае Println внутри себя вызовет функцию и выведет в консоль результат её выполнения
	fmt.Println(findDiscriminant(1, -3.0, 2))
	fmt.Println(findDiscriminant(1, 4.0, 2))
}
