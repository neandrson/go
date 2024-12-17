package main

import "fmt"

// Интерфейс Reader
type Reader interface {
	Read() string
}

// Интерфейс Writer
type Writer interface {
	Write(data string)
}

// Интерфейс ReadWriter объединяет методы Reader и Writer
type ReadWriter interface {
	Reader
	Writer
}

// Структура, которая реализует интерфейс ReadWriter
type TextProcessor struct {
	data string
}

func (tp *TextProcessor) Read() string {
	return tp.data
}

func (tp *TextProcessor) Write(data string) {
	tp.data = data
}

func main() {
	tp := &TextProcessor{}

	// Используем интерфейс ReadWriter
	var rw ReadWriter = tp

	rw.Write("Hello, World!")
	fmt.Println(rw.Read()) // Вывод: Hello, World!
}
