package main

import (
	"sync"
	"testing"
)

func TestConcurrentQueue_EnqueueDequeue(t *testing.T) {
	concurrentQueue := &ConcurrentQueue{}

	testCases := []struct {
		name     string
		elements []interface{}
	}{
		{
			name:     "Test Case 1",
			elements: []interface{}{1, 2, 3},
		},
		{
			name:     "Test Case 2",
			elements: []interface{}{"a", "b", "c"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			for _, element := range tc.elements {
				concurrentQueue.Enqueue(element)
			}

			for _, expected := range tc.elements {
				actual := concurrentQueue.Dequeue()
				if actual != expected {
					t.Errorf("Unexpected element. Got: %v, Expected: %v", actual, expected)
				}
			}

			var wg sync.WaitGroup

			for i := 1; i <= 10000; i++ {
				wg.Add(1)
				go func(item int) {
					concurrentQueue.Enqueue(item)
					wg.Done()
				}(i)
			}
			wg.Wait()

			if len(concurrentQueue.queue) != 10000 {
				t.Errorf("Unexpected len. Got: %v, Expected: %v", len(concurrentQueue.queue), tc.elements)
			}
		})
	}
}
