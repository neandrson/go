package main

import (
	"io"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func WriteData(writer io.Writer, data []byte) error {
	/*bytesWritten, err := writer.Write(data) // Записываем данные
	  if err != nil {
	     // TODO: handle error.
	  }
	  fmt.Printf("Записано %d байт", bytesWritten)*/

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
