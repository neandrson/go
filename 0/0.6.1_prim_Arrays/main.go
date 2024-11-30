package main

import "fmt"

func main() {
	candies := [...]string{"шоколадная", "ванильная", "клубничная", "карамельная", "лимонная"}
	var sum int

	for i := 0; i < len(candies); i++ {
		// съедаем конфету и добавляем её в сумму
		fmt.Println("Я съел конфету вкуса", candies[i])
		sum++
	}

	fmt.Println("Я съел", sum, "конфет")
}
