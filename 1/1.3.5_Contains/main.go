package main

import (
	"fmt"
	"io"
	"strings"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	/*br := bufio.NewReader(r)
	buf := make([]byte, len(seq))
	for {
		n, err := br.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
		for i := 0; i < n; i++ {
			if bytes.Contains(buf[i:], seq) {
				return true, nil
			}
		}
	}
	return false, nil*/

	buf := make([]byte, len(seq))
	pos := 0

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if buf[i] == seq[pos] {
				pos++
				if pos == len(seq) {
					return true, nil
				}
			} else {
				pos = 0
			}
		}
	}

	return false, nil
}

func main() {
	r := strings.NewReader("hello world")
	seq := []byte("world")

	found, err := Contains(r, seq)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sequence found:", found)
	}
}
