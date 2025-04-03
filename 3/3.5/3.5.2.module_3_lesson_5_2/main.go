package main

import (
	"context"
)

type sequenced interface {
	getSequence() int
}

func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return
			default:
				if num.getSequence()%2 == 0 {
					out <- num
				}
			}
		}
	}()

	return out
}
