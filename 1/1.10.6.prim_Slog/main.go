package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	// По умолчанию текстовый формат лога
	/*slog.Debug("Debug message") // 2023/09/12 17:29:42 INFO Debug message

	// Но можно и JSON
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Info message")*/ // {"time":"2023-09-12T17:29:42.64891153+03:00","level":"INFO","msg":"Info message"}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))
	// {"time":"2009-11-10T23:00:00Z","level":"INFO","msg":"finished","req":{"method":"GET","url":"localhost"},"status":200,"duration":1000000000}
}
