package main

/*func doSomething() {
	fmt.Println("hello world")
}*/

func main() {
	// Приемр 1 Go
	//go doSomething()

	// Пример 2 Go
	//time.Sleep(1 * time.Second)

	// Пример 3 Планировщик Go
	//fmt.Println(runtime.NumGoroutine())

	// Ghbvth 4 Каналы
	/*ch := make(chan int)

	go func() {
		ch <- 123 // отправляем значение в канал
	}()

	val := <-ch      // получаем значение из канала
	fmt.Println(val)*/ // выводит 123

	// Пример 5 Дедлок
	ch := make(chan int)
	ch <- 1 // это приведёт к дедлоку, так как нет другой горутины, которая могла бы принять значение
}
