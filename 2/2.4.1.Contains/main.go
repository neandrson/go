package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	buff := make([]byte, len(seq))

	// Начальное чтение данных в буфер
	n, err := r.Read(buff)
	if err != nil && err != io.EOF {
		return false, err
	}
	if n != len(seq) {
		return false, nil
	}

	// Процесс поиска в потоке данных
	/*for {
		if bytes.Equal(seq, buff) {
			return true, nil
		}
		buff = append(buff[1:], 0) // Сдвиг в буфере
		_, err := r.Read(buff[len(buff)-1:])
		if err != nil {
			return false, nil
		}
	}*/
	if err := ctx.Err(); err != nil {
		// time to stop... but why...?
		switch err {
		case context.Canceled:
			// context was cancelled
		case context.DeadlineExceeded:
			// context timed out
		}
	}
}

func main() {
	ctx := context.Background()
	r := bytes.NewReader([]byte{0x61, 0x64, 0x73, 0x61, 0x64, 0x73, 0x61, 0x64, 0x73})
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x64, 0x73})) // true <nil>
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x67, 0x73})) // false <nil>
}
