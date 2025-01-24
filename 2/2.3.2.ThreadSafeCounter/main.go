package main

import (
	"fmt"
	"sync"
)

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

type Counter struct {
	value int // значение счетчика
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	numGoroutines := 100000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	expectedValue := numGoroutines
	actualValue := counter.GetValue()
	if expectedValue != actualValue {
		fmt.Printf("Expected counter value to be %d, but got %d\n", expectedValue, actualValue)
	} else {
		fmt.Printf("counter value but got %d\n", actualValue)
	}
}
