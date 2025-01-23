package main

import (
	"fmt"
	"slices"
)

func Send(ch1, ch2 chan int) {

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			//ch1 <- 1
			//ch1 <- 2
		}
	}()

	//time.Sleep(2 * time.Second)

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i
			//ch2 <- 1
			//ch2 <- 2
		}
	}()

	//time.Sleep(2 * time.Second)
	//close(ch1)
	//close(ch2)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	Send(ch1, ch2)

	vals := []int{}

	for i := 0; i < 3; i++ {
		val := <-ch1
		vals = append(vals, val)
	}

	slices.Sort(vals)

	expected := []int{0, 1, 2}
	if !slices.Equal(vals, expected) {
		fmt.Printf("ch1 values expected: %v, got: %v\n", expected, vals)
	} else {
		fmt.Printf("ch1 values got: %v\n", vals)
	}
}
