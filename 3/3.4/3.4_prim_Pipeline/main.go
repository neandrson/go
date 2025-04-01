package main

import (
	"fmt"
	"sync"
)

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

func merge(cs ...<-chan int) <-chan int {
	// для синхронизации
	var wg sync.WaitGroup
	// объединять будем в этот канал
	out := make(chan int)
	// output - функция, которая копирует данные из
	// входящего канала (пока он не закрыт) в out,
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}
	// добавим по количеству каналов
	wg.Add(len(cs))
	for _, c := range cs {
		// запустим функцию output в отдельном канале
		go output(c)
	}
	// запустим горутину, ожидающую завершения чтения данных из всех входящих
	// каналов. Затем закроем выходящий канал
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Пример 1 -
	// канал в четными числами
	//evens := EvenNumbersGen(1, 2, 3, 4, 5, 6)
	// канал с удвоенными числами
	//out := DoubleNumbers(evens)
	// печатаем каждое число
	//for num := range out {
	//	fmt.Println(num) // 4, 8, 12
	//}

	// Пример 2 -
	// канал в четными числами
	//evens := EvenNumbersGen(1, 2, 3, 4, 5, 6)
	// канал с дважды удвоенными числами
	//out := DoubleNumbers(DoubleNumbers(evens))
	// печатаем каждое число
	//for num := range out {
	//	fmt.Println(num) //8, 16, 24
	//}

	// Пример 3 -
	// канал в четными числами
	evens := EvenNumbersGen(1, 2, 3, 4, 5, 6)
	// Распределим нагрузку. Обе функцию читают из одного канала
	doubled1 := DoubleNumbers(evens)
	doubled2 := DoubleNumbers(evens)
	// Прочитаем значения с объединенного канала
	for n := range merge(doubled1, doubled2) {
		fmt.Println(n) // 4, 8, 12 - могут быть в любом порядке
	}
}
