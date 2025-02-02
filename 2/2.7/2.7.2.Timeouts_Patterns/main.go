package main

import (
	"context"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	data, err := os.ReadFile(path)
	select {
	case <-ctx.Done():
		return
	default:
		if err != nil {
			result <- data
		}
	}
}
