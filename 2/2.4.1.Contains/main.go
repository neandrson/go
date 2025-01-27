package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return true, nil
	}

	buf := make([]byte, 4096)
	window := make([]byte, 0, len(seq))

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:

			n, err := r.Read(buf)
			if err != nil && !errors.Is(err, io.EOF) {
				return false, err
			}

			window = append(window, buf[:n]...)

			if bytes.Contains(window, seq) {
				return true, nil
			}
			if err != nil && err == io.EOF {
				return false, nil
			}

			if len(window) > len(seq) {
				window = window[len(window)-len(seq):]
			}
		}
	}
}

func main() {
	ctx := context.Background()
	r := bytes.NewReader([]byte{0x61, 0x64, 0x73, 0x61, 0x64, 0x73, 0x61, 0x64, 0x73})
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x64, 0x73})) // true <nil>
	fmt.Println(Contains(ctx, r, []byte{0x61, 0x67, 0x73})) // false <nil>
}
