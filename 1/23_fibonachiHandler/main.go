package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := fibonacci(n)
	fmt.Fprintf(w, "%d", result)
}

func StartServer(t time.Duration) {
	http.ListenAndServe(":8080", nil)
}

func main() {
	http.HandleFunc("/fibonacci", FibonacciHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
