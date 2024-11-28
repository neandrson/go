package main

import (
	"log/slog"
	"os"
)

func main() {
	// По умолчанию текстовый формат лога
	slog.Debug("Debug message") // 2023/09/12 17:29:42 INFO Debug message

	// Но можно и JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Info message") // {"time":"2023-09-12T17:29:42.64891153+03:00","level":"INFO","msg":"Info message"}
}
