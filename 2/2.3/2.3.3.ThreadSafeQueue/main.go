package main

import (
	"fmt"
	"sync"
)

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

type ConcurrentQueue struct {
	queue []interface{} // здесь хранить элементы очереди
	mutex sync.Mutex
}

func (c *ConcurrentQueue) Enqueue(element interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.queue = append(c.queue, element)
}

func (c *ConcurrentQueue) Dequeue() interface{} {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.queue) == 0 {
		return nil
	}
	queue := c.queue[0]
	c.queue = c.queue[1:]
	return queue
}

func main() {
	concurrentQueue := &ConcurrentQueue{}

	Cases := []struct {
		name     string
		elements []interface{}
	}{
		{
			name:     "Case 1",
			elements: []interface{}{1, 2, 3},
		},
		{
			name:     "Case 2",
			elements: []interface{}{"a", "b", "c"},
		},
	}

	for _, c := range Cases {
		c := c
		//Run(c.name, func(t *testing.T) {

		for _, element := range c.elements {
			concurrentQueue.Enqueue(element)
		}

		for _, expected := range c.elements {
			actual := concurrentQueue.Dequeue()
			if actual != expected {
				fmt.Printf("Unexpected element. Got: %v, Expected: %v\n", actual, expected)
			} else {
				fmt.Printf("Unexpected element. Got: %v\n", actual)
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
			fmt.Printf("Unexpected len. Got: %v, Expected: %v\n", len(concurrentQueue.queue), 10000)
		} else {
			fmt.Printf("Unexpected len. Got: %v\n", len(concurrentQueue.queue))
		}
		//}
	}
}
