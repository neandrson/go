package main

import (
	"fmt"
	"slices"
	"time"
)

func Send(ch1, ch2 chan int) {
	//ch1 = make(chan int)
	//ch2 = make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i // отправляем значение в канал
			ch2 <- i // отправляем значение в канал
			//time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(1 * time.Second)
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
