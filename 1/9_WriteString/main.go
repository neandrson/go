package main

import (
	"fmt"
	"io"
)

func WriteString(s string, w io.Writer) error {
	_, err := fmt.Fprint(w, s)
	return err
}

func main() {

}
