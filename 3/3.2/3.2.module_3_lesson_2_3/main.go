package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out := make(chan []string)

	go func(out chan []string) {
		defer close(out)
		reader := csv.NewReader(f)
		for {
			data, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return
			}
			/*select {
			case <-data:
				fmt.Println("Получили данные")
				return
			case <-errChan:
				fmt.Println("Получили ошибку")
				return
			}*/
			out <- data
		}
	}(out)

	return out, nil
}

func creatCSV(fileCsv string, data [][]string) error {
	file, err := os.Create(fileCsv)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initialize csv writer
	writer := csv.NewWriter(file)

	defer writer.Flush()

	// write all rows at once
	writer.WriteAll(data)

	// write single row
	extraData := data
	writer.Write(extraData)
}

func main() {

}
