package main

import "fmt"

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		sum := 0

		for {
			select {
			case i := <-in:
				sum += i
			case <-done:
				out <- sum
				return
			}
		}
	}()

	return out
}

func main() {
	in := make(chan int)
	done := make(chan struct{})

	result := DoubleNumbers(done, in)
	for i := 0; i < 10; i++ {
		in <- i
	}

	close(done)
	fmt.Println(<-result)
}
