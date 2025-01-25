package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	b := make([]byte, 1)
	buff := make([]byte, len(seq))

	for {
		_, err := r.Read(b)
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}

		// Сдвиг в буфере для поиска
		for i := 0; i < len(buff)-1; i++ {
			buff[i] = buff[i+1]
		}
		buff[len(buff)-1] = b[0]

		if bytes.Equal(seq, buff) {
			return true, nil
		}
	}
}

func main() {
	ctx := context.Background()
	r := bytes.NewReader([]byte{0x61, 0x64, 0x73, 0x61, 0x64, 0x73, 0x61, 0x64, 0x73})
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x64, 0x73})) // true <nil>
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x67, 0x73})) // false <nil>
}
