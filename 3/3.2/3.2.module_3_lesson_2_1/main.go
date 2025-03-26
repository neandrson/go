package main

import "fmt"

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	output := make(chan int)
	var x int

	go func(output chan int) {
		defer close(output)
		for {
			select { // Оператор select
			case <-in: // Ждет, когда проснется гофер
				x *= <-in
				output <- x
				fmt.Println(x)
			case <-done: // Ждет окончания времени
				return
			}
		}

	}(output)

	return output
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
