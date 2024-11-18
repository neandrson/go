package main

import (
	"fmt"
	"strings"
)

// Манипуляции со строками

// Интерфейс io.Writer выглядит так:
// type Writer interface {
// Write(p []byte) (n int, err error)
// }

// Чтобы реализовать интерфейс, все методы, указанные интерфейсом,
// должны быть определены для типа структуры

type UpperWriter struct {
	UpperString string
}

// Функции, связанные с типом UpperWriter

// Реализуем интерфейс io.Writer

func (uw *UpperWriter) Write(p []byte) (n int, err error) {
	// Построим строку из байтового слайса
	str := string(p[:])
	// Преобразуем строку в верхний регистр
	uw.UpperString = strings.ToUpper(str)
	// Вернём длину строки и отсутствие ошибки
	return len(uw.UpperString), nil
}

func main() {
	// Определим массив байт со словом 'Hello'
	byteArray := []byte{'H', 'e', 'l', 'l', 'o'}
	// Определим объект UpperWriter
	var uw UpperWriter
	fmt.Println("UpperWriter test")
	// Запишем массив байт в объект при помощи метода Write
	n, _ := uw.Write(byteArray)
	// Выведем исходник и результат преобразования
	fmt.Printf("'%s' => '%s'[len: %d]\n", string(byteArray[:]), uw.UpperString, n)
}
