package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	//"golang.org/x/exp/rand"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	defer s.mux.Unlock()
	data, _ := s.m[key] // теперь доступ к мапе внутри критической секции
	return data
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.m[key] = value
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{})}
}

func main() {
	runtime.GOMAXPROCS(10)

	concurrentMap := NewSafeMap()

	var wg sync.WaitGroup

	wg.Add(2000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			concurrentMap.Set(fmt.Sprintf("%d", rand.Intn(100)), rand.Intn(100))
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			concurrentMap.Get(fmt.Sprintf("%d", rand.Intn(100)))
		}()
	}

	wg.Wait()

	var wg2 sync.WaitGroup

	wg2.Add(100)

	for i := 0; i < 100; i++ {
		concurrentMap.Set(fmt.Sprintf("%d", i), i)
	}

	for i := 0; i < 100; i++ {
		i := i
		go func() {
			defer wg2.Done()
			val := concurrentMap.Get(fmt.Sprintf("%d", i))

			if val.(int) != i {
				fmt.Printf("Expected %d, got %d\n", i, val)
			} else {
				fmt.Printf("got %d\n", val)
			}
		}()
	}
	wg2.Wait()
}
