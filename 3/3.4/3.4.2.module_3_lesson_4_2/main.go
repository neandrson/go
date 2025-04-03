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

	/*out := make(chan int) // канал для записи выходных данных
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		return nil
	}
	go func() {

		defer file.Close()

		data := bufio.NewReader(file)
		//lineNumber := 1
		for data.Read() {
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
	return out // вернём канал*/

	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	out := make(chan int)
	r := bufio.NewReader(file)
	go func() {

		defer close(out)
		//data := make([]byte, 64)

		l, s, e := r.ReadLine()
		for s && e == nil {
			fmt.Println(string(l))
			//s, e = Readln(r)

		}

	}()
	return out
}

func main() {

}
