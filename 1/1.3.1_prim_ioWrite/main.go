package main

import (
	"io"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func WriteData(writer io.Writer, data []byte) error {
	// Пример 1
	/*bytesWritten, err := writer.Write(data) // Записываем данные
	  if err != nil {
	     // TODO: handle error.
	  }
	  fmt.Printf("Записано %d байт", bytesWritten)*/

	// Пример 2
	file, err := os.Create("output.txt") // Создаём файл
	if err != nil {
		// TODO: handle error.
	}
	defer file.Close() // Закроем файл после записи
	data := []byte("Hello, World!")
	if err = WriteData(file, data); err != nil { // Запишем данные в файл
		// TODO: handle error.
	}
}

// Пример 2
type myWriter struct {
	content []byte // Сюда будем записывать данные
}

// Реализуем интерфейс io.Writer
func (w *myWriter) Write(buf []byte) (int, error) {
	w.content = append(w.content, buf...) // Запишем данные в слайс
	return len(buf), nil                  // Вернём, сколько данных записано
}

// Для удобного представления реализуем интерфейс Stringer
func (w *myWriter) String() string {
	return string(w.content)
}

/*func WriteData(writer io.Writer, data []byte) error {
	w := &myWriter{}
	WriteData(w, []byte("привет\n")) // Запишем данные в слайс
	fmt.Println(w.String()) // "привет\n"
}*/
