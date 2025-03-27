package main

import "fmt"

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string, 3) // канал для записи выходных данных
	go func() {                 // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for i := range valueStream {
			select {
			case <-done:
				return // после закрытия канала done - выходим
			default:
				out <- fmt.Sprint(i) // запишем в канал
			}
		}
	}()
	return out
}

func main() {
	//type T any{}

	done := make(chan struct{})
	in := make(chan int)
	in <- 1
	out := ToString(done, in)
	fmt.Println(out)
}
