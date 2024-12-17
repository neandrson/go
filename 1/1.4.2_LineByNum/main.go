package main

import (
	"bufio"
	"log"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	// Открытие файла на чтение
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Создание сканера для чтения файла построчно
	scanner := bufio.NewScanner(file)
	// Счётчик текущей строки
	currentLineNum := 0
	// Последняя найденная строка
	var lineText string
	// Чтение файла до достижения нужной строки или конца файла
	for scanner.Scan() {
		if currentLineNum == lineNum {
			lineText = scanner.Text()
			break
		}
		currentLineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lineText
}

func main() {

}
