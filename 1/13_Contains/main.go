package main

import (
	"bufio"
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	br := bufio.NewReader(r)
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
	return false, nil
}

func main() {

}
