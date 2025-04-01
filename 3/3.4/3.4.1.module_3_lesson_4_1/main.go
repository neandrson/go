package main

import "fmt"

func StringsGen(lines ...string) <-chan string {
	out := make(chan string) // канал для записи выходных данных
	go func() {              // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for _, line := range lines {
			//if num%2 == 0 {
			out <- line // запишем в канал
			//}
		}
	}()
	return out // вернём канал
}

func main() {
	// канал текстовый
	linens := StringsGen("1", "2", "3", "4", "5", "6")
	for line := range linens {
		fmt.Println(line)
	}
}
