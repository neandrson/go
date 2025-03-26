package main

import "fmt"

type any interface {
	int | float64 | string
}

func ToString[T any](done <-chan struct{}, valueStream <-chan T) <-chan string {
	out := make(chan string) // канал для записи выходных данных
	go func() {              // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for i := range valueStream {
			select {
			case <-done:
				return // после закрытия канала done - выходим
			default:
				s := fmt.Sprintf("%d", i)
				out <- s // запишем в канал
			}
		}
	}()
	return out
}

func main() {

}
