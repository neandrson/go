package main

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	out := make(chan []string)

	go func() {
		defer f.Close()
		defer close(out)

		reader := csv.NewReader(f)
		for {
			data, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				//log.Fatal(err)
				return
			}
			out <- data
		}
	}()

	return out, nil
}

func creatCSV(fileCsv string, data [][]string) error {
	file, err := os.Create(fileCsv)
	if err != nil {
		//log.Fatal(err)
		return err
	}
	defer file.Close()

	// initialize csv writer
	writer := csv.NewWriter(file)

	defer writer.Flush()

	// write all rows at once
	writer.WriteAll(data)

	// write single row
	//extraData := data
	//writer.Write(extraData)
	return nil
}

func main() {

}
