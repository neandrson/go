package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	var err error // Ошибка будет тут
	var rlen int  // Кол-во прочитанных байт
	// Породим буфер размером n байт
	buf := make([]byte, n)
	rlen, err = r.Read(buf)
	if err == nil {
		if rlen != 0 {
			out_slice := buf[:rlen]
			_, err = w.Write(out_slice)
		}
	} else {
		if err == io.EOF {
			err = nil
		}
	}
	return err
}

/*
func Copy(r io.Reader, w io.Writer, n uint) error {
    data := make( []byte, n)
    _, err := r.Read(data)
    if err != nil && err != io.EOF {
        return err
    }
    _, err = w.Write(data)
    if err != nil {
        return err
    }
    return nil
} */

func main() {
	// Пример использования функции Copy
	r := strings.NewReader("Hello, World!")
	w := os.Stdout
	n := uint(5)
	err := Copy(r, w, n)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
