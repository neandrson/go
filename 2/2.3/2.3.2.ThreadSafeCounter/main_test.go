package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
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
		t.Errorf("Expected counter value to be %d, but got %d", expectedValue, actualValue)
	}
}
