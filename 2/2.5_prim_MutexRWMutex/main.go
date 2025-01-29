package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// функция генерирует случайное число в интервале [0, 100)
func random() int {
	const max int = 100
	return rand.Intn(max)
}

func main() {
	const size int = 10
	results := []int{}
	mx := &sync.Mutex{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		// Пример 1
		//results = append(results, random())

		// Пример 2
		go func() {

			// Пример 3
			mx.Lock()
			defer mx.Unlock()

			results = append(results, random())
		}()
	}

	// Пример 2
	time.Sleep(time.Second)

	// Пример 3
	// вызван Lock, потому что здесь тоже обращаемся к results
	mx.Lock()
	defer mx.Unlock()

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}
}
