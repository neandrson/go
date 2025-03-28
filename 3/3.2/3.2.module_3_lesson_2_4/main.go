package main

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T) // канал для записи выходных данных
	go func() {          // запускаем в отдельной горутине
		defer close(out1) // закроем канал, когда больше нет данных
		defer close(out2)
		for i := range in {
			select {
			case <-done:
				return // после закрытия канала done - выходим
			default:
				out1 <- i // запишем в канал
				out2 <- i
			}
		}
	}()
	return out1, out2
}

func main() {

}
