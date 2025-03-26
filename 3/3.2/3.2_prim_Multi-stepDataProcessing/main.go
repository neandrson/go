package main

import (
	"crypto/rand"
	"fmt"
)

// Пример 1 -
/*func EvenNumbersGen(done <-chan struct{}, numbers ...int) <-chan int {
	out := make(chan int) // канал для записи выходных данных
	go func() {           // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for _, num := range numbers {
			select {
			case <-done:
				return // после закрытия канала done - выходим
			default:
				if num%2 == 0 {
					out <- num // запишем в канал
				}
			}
		}
	}()
	return out // вернём канал
}

func main() {
	done := make(chan struct{})
	defer close(done) // закроем канал при завершении main
	// канал c четными числами
	evens := EvenNumbersGen(done, 1, 2, 3, 4, 5, 6)
	fmt.Println(<-evens) // прочитаем первое значение
	// ...
}*/

// Пример 2 - Генераторы — это любая функция, которая преобразует набор дискретных значений в поток данных в канале.
// Один из простых генераторов — repeat
/*func Repeat[T any](done <-chan struct{}, values ...T) <-chan T { // используем дженерики для передачи любого значения
	out := make(chan T)
	go func() {
		defer close(out)
		for { // бесконечный цикл
			for _, v := range values {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()
	return out
}

func Take[T any](done <-chan struct{},valueStream <-chan T,	num int) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ { // ограниченное число значений
			select {
			case <-done:
				return
			case out <- <-valueStream:
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)
	out := Take(done, Repeat(done, 2), 4)
	for v := range out {
		fmt.Println(v) // 2 2 2 2
	}
}*/

// Пример 3 -
func RepeatFunc(done <-chan struct{}, fn func() int) <-chan int { // функция, которая будет вызываться бесконечно
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- fn(): // запишем в канал результат работы функции
			}
		}
	}()
	return out
}

func Take[T any](done <-chan struct{}, valueStream <-chan T, num int) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ { // ограниченное число значений
			select {
			case <-done:
				return
			case out <- <-valueStream:
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)
	// функция генерации радномного числа
	randFunc, _ := func() int, error { return rand.Int() }
	out := Take(done, RepeatFunc(done, randFunc), 5)
	for v := range out {
		fmt.Println(v) // здесь будут пять случайных чисел
	}
}
