package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Work struct {
	file    string         // путь к файлу
	pattern *regexp.Regexp // регулярное выражение для поиска в файле
}

// Данные о файлах в директории будут передаваться в канал jobs на первом этапе обработки:
func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	//	jobs := make(chan Work) // канал для записи информации о файлах

	// Пример 3 - Если есть необходимость быстро выполнить сканирование файлов в директории, для этой цели можно использовать буферизированные каналы
	// сюда можно записать 10 элементов без блокировки
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

// Сама функция поиска будет получать файлы из канала jobs:
func worker(jobs <-chan Work) {
	// получаем информацию об очередном файле из канала
	for work := range jobs {
		// открываем файл
		f, err := os.Open(work.file)
		if err != nil {
			fmt.Println(err)
			continue // пропустим ошибки для упрощения
		}
		// scanner для чтения построчно
		scn := bufio.NewScanner(f)
		lineNumber := 1
		// читаем файл
		for scn.Scan() {
			// поиск в каждой строке
			result := work.pattern.Find(scn.Bytes())
			// если нашли - выведем результат на экран
			if len(result) > 0 {
				fmt.Printf("%s#%d: %s\n", work.file,
					lineNumber, string(result))
			}
			lineNumber++
		}
		f.Close()
	}
}

func main() {
	// регулярное выражение для поиска
	pattern := regexp.MustCompile(os.Args[2])
	// сканирование директории
	jobs := FileNameGen(os.Args[1], pattern)
	// поиск внутри каждого файла
	//worker(jobs)

	// Пример 2 -
	wg := sync.WaitGroup{}   //  для ожидания всех обработчиков
	for i := 0; i < 3; i++ { // ограничим размер пула до трёх
		wg.Add(1)
		go func() {
			defer wg.Done() // отметим, что обработчик завершил работу
			worker(jobs)
		}()
	}
	wg.Wait() // дождёмся окончания работы всех обработчиков
}
