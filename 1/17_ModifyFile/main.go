package main

import (
	"io"
	"os"
)

func ModifyFile(filename string, pos int, val string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Seek(int64(pos), io.SeekStart)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(val))
	if err != nil {
		return err
	}
	return nil
}

func main() {

}
