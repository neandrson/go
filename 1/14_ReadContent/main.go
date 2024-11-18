package main

import (
	"io/ioutil"
	"log"
)

func ReadContent(filename string) string {
	// Чтение содержимого файла
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(content)
}

func main() {
	filename := "file.txt"

	fileContent := ReadContent(filename)
	log.Printf("Содержимое файла %s:\n%s", filename, fileContent)
}
