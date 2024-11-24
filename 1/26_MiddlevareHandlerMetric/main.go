package main

import (
	"fmt"
	"net/http"
	"time"
)

var requestCount int

func Metrics(next http.HandlerFunc) http.HandlerFunc {
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
	fibonacciHandler := http.HandlerFunc(fibonacciHandler)
	metricsHandler := http.HandlerFunc(metricsHandler)
	http.Handle("/", Metrics(fibonacciHandler))
	http.Handle("/metrics", metricsHandler)
	http.ListenAndServe(":8080", nil)
}
