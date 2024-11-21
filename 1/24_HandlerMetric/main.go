package main

import (
	"fmt"
	"net/http"
	"time"
)

/*var current int = 0
var next int = 1
var requestCount int = 0

func Next() int {
	result := current
	current, next = next, current+next
	return result
}

// func Metrics(next http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     next.ServeHTTP(w, r)
//     requestCount++
//   })
// }

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	number := Next()

	fmt.Fprintf(w, "%d", number)
}

func StartServer(shutdownAfter time.Duration) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", FibonacciHandler)
	mux.HandleFunc("/metrics", MetricsHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		time.Sleep(shutdownAfter)
		server.Close()
	}()

	fmt.Println("Сервер запущен на :8080")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}*/

var requestCount int

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("rpc_duration_milliseconds_count %d\n", duration)
		requestCount++
	})
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	n := fibonacci(requestCount)
	fmt.Fprintf(w, "%d", n)
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
}

func main() {
	fibonacciHandler := http.HandlerFunc(FibonacciHandler)
	metricsHandler := http.HandlerFunc(MetricsHandler)

	http.Handle("/", Metrics(fibonacciHandler))
	http.Handle("/metrics", metricsHandler)

	http.ListenAndServe(":8080", nil)
}
