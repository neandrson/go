package main

import (
	"fmt"
	"time"
)

// Пример 4
/*type SafeSlice struct {
	results []int
	mx      *sync.Mutex
}

// добавляем к слайсу элемент item
func (s *SafeSlice) Append(item int) {
	// вызван Lock, поэтому только одна горутина за раз может получить доступ к
	// слайсу
	s.mx.Lock()
	defer s.mx.Unlock()
	s.results = append(s.results, random())
}

// получаем элемент слайса по индексу
func (s *SafeSlice) Get(index int) int {
	// вызван Lock, поэтому только одна горутина за раз может получить доступ к
	// слайсу
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.results[index]
}

// создаёт экземпляр нашей структуры, для удобства пользователя
func New() *SafeSlice {
	return &SafeSlice{
		mx:      &sync.Mutex{},
		results: []int{},
	}
}

// Пример 1
// функция генерирует случайное число в интервале [0, 100)
func random() int {
	const max int = 100
	return rand.Intn(max)
}*/

// Пример 1
func main() {
	// Пример 1
	/*const size int = 10
	results := []int{}
	mx := &sync.Mutex{}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {

		//results = append(results, random())

		// Пример 2
		go func() {

			// Пример 3
			mx.Lock()
			defer mx.Unlock()

			// Пример 2
			results = append(results, random())
		}()
	}

	// Пример 2
	time.Sleep(time.Second)

	// Пример 3
	// вызван Lock, потому что здесь тоже обращаемся к results
	mx.Lock()
	defer mx.Unlock()

	// Пример 1
	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(results[i])
	}*/

	// Пример 4
	/*safeSlice := New()
	const size int = 10
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func() {
			safeSlice.Append(random())
		}()
	}
	time.Sleep(time.Second)

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(safeSlice.Get(i))
	}*/

	// Пример 5 Каналы
	ch := make(chan struct{})
	// горутина, которая асинхронно производит вычисления
	go func() {
		fmt.Println("начинаем вычисления...")
		// имитируем длинные вычисления
		time.Sleep(time.Second)
		fmt.Println("заканчиваем вычисления ...")
		// закрываем канал, чтобы получить сообщения
		close(ch)
	}()

	// программа блокируется
	<-ch
	fmt.Println("завершаем программу")
}
