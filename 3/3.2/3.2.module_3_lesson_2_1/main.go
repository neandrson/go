package main

import "fmt"

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int) // канал для записи выходных данных
	go func() {           // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for i := range in {
			select {
			case <-done:
				return // после закрытия канала done - выходим
			default:
				out <- 2 * i // запишем в канал
			}
		}
	}()
	return out
}

func main() {
	in := make(chan int)
	done := make(chan struct{})
	//num := []int{1, 2, 3}
	result := make(chan int)
	for i := 0; i < 10; i++ {
		in <- i
		result = DoubleNumbers(done, in)
	}
	/*for i := range num {
		in <- i
	}*/

	close(done)
	fmt.Println(<-result)
}
