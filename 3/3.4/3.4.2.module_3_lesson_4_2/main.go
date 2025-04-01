package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func NumbersGen(filename string) <-chan int {
	out := make(chan int) // канал для записи выходных данных
	file, err := os.Open("hello.txt")
	if err != nil {
		return nil
	}
	defer file.Close()
	data := make([]byte, 0)
	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		num := fmt.Sprint(n)
		_, err = strconv.Atoi(num)
		if err == nil {
			continue
		}
		out <- n
	}
	return out // вернём канал
}

func main() {

}
