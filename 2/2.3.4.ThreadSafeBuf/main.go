package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var Buf []int

func Write(num int) {
	mu.Lock()
	defer mu.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mu.Lock()
	defer mu.Unlock()
	/*if len(c.queue) == 0 {
		return nil
	}*/
	buf := Buf[0]
	//Buf = Buf[1:]
	return buf
}

func main() {
	Buf := []int{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			Write(i)
		}()
	}

	wg.Wait()

	if len(Buf) != 1000 {
		fmt.Printf("Unexpected len. Got: %d, Expected: %d\n", len(Buf), 1000)
	} else {
		fmt.Printf("Unexpected len. Got: %d\n", len(Buf))
	}

	var results []int

	var resultMu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			res := Consume()

			resultMu.Lock()
			results = append(results, res)
			resultMu.Unlock()
		}()
	}
	wg.Wait()

	if len(Buf) != 0 {
		fmt.Printf("Unexpected len. Got: %d, Expected: %d\n", len(Buf), 0)
	} else {
		fmt.Printf("Unexpected len. Got: %d\n", len(Buf))
	}
	if len(results) != 1000 {
		fmt.Printf("Unexpected results len. Got: %d, Expected: %d", len(results), 1000)
	} else {
		fmt.Printf("Unexpected results len. Got: %d\n", len(results))
	}
}
