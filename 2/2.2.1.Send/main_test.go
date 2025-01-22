package main

import "testing"

func TestSend(t *testing.T) {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go Send(ch, i)
		val := <-ch

		if val != i {
			t.Fatalf("Expected to receive: %v, got: %v", i, val)
		}
	}
}
