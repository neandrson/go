package main

import "io"

type WordCounter struct {
	wordsCount map[string]int // здесь должен быть список слов с указанием количества повторений во всех файлах.
	// можно добавлять другие поля
}

type CounterWorker interface {
	ProcessFiles(files ...string) error // для запуска обработки файлов
	ProcessReader(r io.Reader) error    // для подсчёта слов в одном файле
}

func main() {

}
