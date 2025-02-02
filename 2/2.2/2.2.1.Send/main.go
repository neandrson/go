package main

import (
	"fmt"
	"time"
)

func Send(ch chan int, num int) {
	ch <- num
}

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go Send(ch, i)
		time.Sleep(1 * time.Second)
		val := <-ch

		if val != i {
			fmt.Printf("Expected to receive: %v, got: %v", i, val)
		} else {
			fmt.Printf("Expected got: %v\n", val)
		}
	}
}
