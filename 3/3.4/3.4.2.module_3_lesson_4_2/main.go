package main

import (
	"bufio"
	"fmt"
	"os"
)

func NumbersGen(filename string) <-chan int {
	/*out := make(chan int) // канал для записи выходных данных
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer close(out) // закроем канал, когда больше нет данных
		for {
			_, err = strconv.Atoi(string(bytes[:]))
			if err == nil {
				out <- n // запишем в канал
			}
		}
	}()
	return out // вернём канал*/

	out := make(chan int) // канал для записи выходных данных
	go func() {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Ошибка открытия файла: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		//lineNumber := 1
		for scanner.Scan() {
			var x interface{}
			switch v := x.(type) {
			case int:
				out <- v
			default:
				continue
			}

			//lineNumber++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Ошибка сканирования: %v\n", err)
		}
	}()
	return out // вернём канал
}

func main() {

}
