package main

import "fmt"

func EvenNumbersGen(numbers ...int) <-chan int {
	out := make(chan int) // канал для записи выходных данных
	go func() {           // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for _, num := range numbers {
			if num%2 == 0 {
				out <- num // запишем в канал
			}
		}
	}()
	return out // вернём канал
}

func DoubleNumbers(in <-chan int) <-chan int {
	out := make(chan int) // выходной канал
	go func() {
		defer close(out) // закроем канал
		for num := range in {
			out <- num * 2
		}
	}()
	return out // вернём канал
}

func main() {
	// канал в четными числами
	evens := EvenNumbersGen(1, 2, 3, 4, 5, 6)
	// канал с удвоенными числами
	out := DoubleNumbers(evens)
	// печатаем каждое число
	for num := range out {
		fmt.Println(num) // 4, 8, 12
	}
}
