package main

import (
	"fmt"
	"net/http"
	"time"
)

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
	requestCount++

}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Println(w, requestCount, r)
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d\n", requestCount)

}

func main() {
	fibonacciHandler := http.HandlerFunc(FibonacciHandler)
	metricsHandler := http.HandlerFunc(MetricsHandler)

	http.Handle("/", Metrics(fibonacciHandler))
	http.Handle("/metrics", metricsHandler)

	http.ListenAndServe(":8080", nil)
}
