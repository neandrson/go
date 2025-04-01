package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
	"sync"
)

type WordCounter struct {
	wordsCount map[string]int // здесь должен быть список слов с указанием количества повторений во всех файлах.
	// можно добавлять другие поля
}

type CounterWorker interface {
	ProcessFiles(files ...string) error // для запуска обработки файлов
	ProcessReader(r io.Reader) error    // для подсчёта слов в одном файле
}

func (wc *WordCounter) ProcessFiles(files ...string) error {
	var errs error
	wg := sync.WaitGroup{}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			errs = errors.Join(errs, err)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer f.Close()
			wc.ProcessReader(f)
		}()
	}

	wg.Wait()
	return errs
}

func (wc *WordCounter) ProcessReader(r io.Reader) error {
	scnr := bufio.NewScanner(r)
	scnr.Split(bufio.ScanWords)
	for scnr.Scan() {
		if scnr.Err() != nil {
			break
		}
		wc.wordsCount[strings.ToLower(scnr.Text())]++
	}

	return scnr.Err()
}

func main() {

}
