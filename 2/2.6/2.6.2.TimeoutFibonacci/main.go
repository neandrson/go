package main

import (
	"errors"
	"fmt"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int)
	go func() {
		fib1, fib2 := 0, 1
		for n > 0 {
			fib1, fib2 = fib2, fib1+fib2
			n--
		}
		c <- fib1
	}()

	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	Cases := []struct {
		input    int
		timeout  time.Duration
		expected int
	}{
		{10, 1 * time.Second, 55},
		{40, 5 * time.Second, 102334155},
	}

	//iterating over every testcases
	for _, c := range Cases {
		result, err := TimeoutFibonacci(c.input, c.timeout)
		if err != nil {
			fmt.Printf("timeoutFibonacci(%d, %v) returned an error: %v\n", c.input, c.timeout, err)
		} else {
			fmt.Printf("timeoutFibonacci(%d, %v)\n", c.input, c.timeout)
		}
		if result != c.expected {
			fmt.Printf("timeoutFibonacci(%d, %v) = %d; want %d\n", c.input, c.timeout, result, c.expected)
		} else {
			fmt.Printf("timeoutFibonacci(%d, %v) = %d\n", c.input, c.timeout, result)
		}
	}
}
