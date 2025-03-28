package main

import (
	"io/fs"
	"path/filepath"
	"regexp"
)

type Work struct {
	FilePath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work, 10)
	go func() {
		defer close(jobs) // закроем канал после обхода всех файлов
		// функция для перебора файлов в директории
		filepath.Walk(dir, func(path string, d fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// пропускаем вложенные директории
			if !d.IsDir() {
				// запишем в канал файл, которые нужно обработать на следующем этапе
				jobs <- Work{file: path, pattern: pattern}
			}
			return nil
		})
	}()
	return jobs
}

func compress(jobs <-chan Work) {

}

func main() {

}
