package main

import (
	"fmt"
	"time"
)

func Receive(ch chan int) int {
	num := <-ch
	return num
}

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func() { ch <- i }()
		time.Sleep(1 * time.Second)

		val := Receive(ch)

		if val != i {
			fmt.Printf("Expected to receive: %v, got: %v\n", i, val)
		} else {
			fmt.Printf("Expected got: %v\n", val)
		}
	}
}
