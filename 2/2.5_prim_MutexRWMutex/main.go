package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Пример 4
type SafeSlice struct {
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
}

// слушатель, которая ждёт сигнала о том, что можно начинать обработку данных
func listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	// отправляем сигнал слушателям
	c.Broadcast()
	c.L.Unlock()
}

// Пример 1
func main() {
	// Пример 1
	//const size int = 10
	//results := []int{}
	/*mx := &sync.Mutex{}
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
	/*ch := make(chan struct{})
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
	fmt.Println("завершаем программу")*/

	// Пример 6 Каналы - итог
	//safeSlice := New()
	//const size int = 10
	// объявляем слайс каналов
	/*channels := make([]chan struct{}, size)
	// создаём каналы функцией make
	for i := range channels {
		channels[i] = make(chan struct{})
	}
	// заполняем слайс случайными числами
	for i := 0; i < size; i++ {
		go func(i int) {
			safeSlice.Append(random())
			// закрываем канал после выполнения задачи
			close(channels[i])
		}(i)
	}
	// ждём, пока не получим сообщения из всех каналов
	for i := 0; i < size; i++ {
		<-channels[i]
	}

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(safeSlice.Get(i))
	}*/

	// Пример 7 - WaitGroup
	// создаём экземпляр WaitGroup
	//wg := &sync.WaitGroup{}

	// заполняем слайс случайными числами
	/*for i := 0; i < size; i++ {
		// добавляем в группу один элемент
		wg.Add(1)
		go func() {
			// удаляем один элемент из группы
			defer wg.Done()
			safeSlice.Append(random())
		}()
	}
	// ждём выполнения всех горутин группы
	wg.Wait()

	// поэлементно выводим слайс на экран
	for i := 0; i < size; i++ {
		fmt.Println(safeSlice.Get(i))
	}*/

	// Пример 8
	// функция, которая имитирует инициализацию ресурсов
	/*initializeResources := func() {
		time.Sleep(time.Second)
		fmt.Println("Only once initialize something")
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			initializeResources()
		}()
	}
	wg.Wait()*/

	// Пример 9 - Cond
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})

	go listen("слушатель 1", data, cond)
	go listen("слушатель 2", data, cond)

	go broadcast("источник", data, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
