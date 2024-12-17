package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startPos int) error {

	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	_, err = inputFile.Seek(int64(startPos), 0)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}
	return nil
}

func main() {

}
