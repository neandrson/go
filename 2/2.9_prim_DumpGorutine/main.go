package main

import (
	"fmt"
	"sync"
)

// Дамп горутин — инструмент, который позволяет увидеть текущее состояние всех горутин в программе и их стеки вызовов
/*func main() {
	// создание файла для записи дампа горутин
	f, err := os.Create("goroutine_dump.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer f.Close()

	// запись дампа горутин в файл каждую секунду
	go func() {
		for {
			// получение дампа горутин
			pprof.Lookup("goroutine").WriteTo(f, 1)
			time.Sleep(time.Second)
		}
	}()

	// ваш код, в котором создаются горутины
	// ...

	// пример бесконечного цикла для демонстрации
	for {
		// здесь может быть ваша основная логика программы
		time.Sleep(1 * time.Second)
	}
}*/

// Race Detector (детектор гонок) — инструмент, который выявляет гонки данных в параллельных программах
/*func main() {
	var wg sync.WaitGroup
	var sharedData int
	iterations := 100

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			sharedData++ // несинхронизированный доступ к общим данным
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of sharedData:", sharedData)
}*/

// Исправлен
func main() {
	var wg sync.WaitGroup
	var sharedData int
	var mu sync.Mutex // создаём мьютекс

	iterations := 100

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // блокировка мьютекса перед доступом к общим данным
			sharedData++
			mu.Unlock() // разблокировка мьютекса после завершения доступа
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of sharedData:", sharedData)
}
