package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}

	wg.Wait()

	if value := counter.GetValue(); value != 500 {
		t.Errorf("Expected counter value: 500, got: %v", value)
	}
}
