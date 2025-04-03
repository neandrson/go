package main

import (
	"context"
	"fmt"
	"sync"
)

type sequenced interface {
	getSequence() int // порядковый номер части сообщения
}

type Num struct {
	int
}

type fanInRecord[T sequenced] struct {
	index int           // порядковый номер горутины, из которой получено сообщение
	data  T             // непосредственно данные
	pause chan struct{} // канал для синхронизации
}

// getSequence реализует интерфейс sequenced.
func (n Num) getSequence() int {
	// для упрощения примера будем считать
	// само число как порядковый номер части
	return n.int
}

// FanIn - объединяет данные из нескольких каналов в один.
func FanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	outputCh := make(chan int)    // выходной канал
	wg := sync.WaitGroup{}        // для ожидания завершения
	for _, ch := range channels { // переберём все каналы
		wg.Add(1)
		// для каждого канала вызовем функцию в отдельной горутине
		go func(input <-chan int) {
			defer wg.Done() // отметим, что функция завершилась
			for {           // цикл для получения данных из входных каналов
				select {
				case data, ok := <-input: // получим данные из канала
					if !ok {
						return // данных больше нет - выходим
					}
					outputCh <- data // запишем данные в выходной канал
				case <-ctx.Done(): // если нужно завершить - выходим
					return
				}
			}
		}(ch) // передадим входной канал в функцию
	}
	go func() {
		wg.Wait()       // дождёмся завершения обработки всех каналов
		close(outputCh) // закроем выходной канал
	}()

	return outputCh // вернём канал
}

// генератор чисел
// func EvenNumbersGen(ctx context.Context, numbers ...int) <-chan int {
func EvenNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T) // канал для записи выходных данных
	go func() {         // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return // выходим
			default:
				if num%2 == 0 {
					out <- num // запишем в канал
				}
			}
		}
	}()
	return out // вернём канал
}

// func OddNumbersGen(ctx context.Context, numbers ...int) <-chan int {
func OddNumbersGen[T sequenced](ctx context.Context, numbers ...T) <-chan T {
	out := make(chan T) // канал для записи выходных данных
	go func() {         // запускаем в отдельной горутине
		defer close(out) // закроем канал, когда больше нет данных
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return // выходим
			default:
				if num%2 == 0 {
					out <- num // запишем в канал
				}
			}
		}
	}()
	return out // вернём канал
}

// inTemp - записывает данные из каналов в один выходной с ожиданием.
func inTemp[T sequenced](ctx context.Context, channels ...<-chan T) <-chan fanInRecord[T] {
	// канал для ожидания
	fanInCh := make(chan fanInRecord[T])
	// для синхронизации
	wg := sync.WaitGroup{}
	// перебор всех входных каналов
	for i := range channels {
		wg.Add(1)
		// запустим горутину для получения данных из канала
		go func(index int) {
			defer wg.Done()
			// канал для синхронизации
			pauseCh := make(chan struct{})
			// цикл для получения данных из канала
			for {
				select {
				// получим данные из канала
				case data, ok := <-channels[index]:
					if !ok {
						return // канал закрыт - выходим
					}
					// положим во временный канал вместе с индексом
					fanInCh <- fanInRecord[T]{
						// индекс канала, откуда пришли данные
						index: index,
						// данные из канала
						data: data,
						// канал для синхронизации
						pause: pauseCh,
					}
				case <-ctx.Done():
					return
				}
				// ждём, пока в канал pause не будет передан сигнал
				// о получении очередного элемента из канала
				select {
				case <-pauseCh:
				// сняли с паузы
				// продолжим обработку данных из входного канала
				case <-ctx.Done():
					return
				}
			}
		}(i)
	}
	go func() {
		// ожидаем завершения
		wg.Wait()
		close(fanInCh)
	}()
	// вернём канал с неотсортированными элементами
	return fanInCh
}

// функционал получения данных из временного канала и, непосредственно, синхронизации для упорядочивания элементов
func processTempCh[T sequenced](ctx context.Context, inputChannelsNum int, // количество входных каналов
	fanInCh <-chan fanInRecord[T], // временный канал с данными
) <-chan T {
	// выходной канал с упорядоченными данными
	outputCh := make(chan T)
	go func() {
		defer close(outputCh)
		// порядковый номер очередного элемента
		expected := 0
		// буфер для ожидания элементов по количеству входных каналов
		queuedData := make([]*fanInRecord[T], inputChannelsNum)
		for in := range fanInCh {
			// если получили элемент с номером, который ожидаем
			if in.data.getSequence() == expected {
				select {
				// запишем элемент в выходной канал
				case outputCh <- in.data:
					// снимем с паузы исходный канал
					// для продолжения обработки из входного канала
					in.pause <- struct{}{}
					// инкремент номера очередного элемента
					expected++
					// здесь нужно реализовать запись в выходной канал
					// из буфера queuedData (задача для домашней работы)
				case <-ctx.Done():
					return
				}
			} else {
				// если НЕ получили элемент с номером, который ожидаем
				// запишем элемент в буфер
				in := in
				queuedData[in.index] = &in
			}
		}
	}()
	return outputCh
}

func main() {
	ctx := context.WithoutCancel()
	// исходный слайс
	//nums := []int{0, 1, 2, 3, 4, 5, 6}
	// канал с чётными числами
	//inputCh1 := EvenNumbersGen(ctx, nums...)
	// канал с нечётными числами
	//inputCh2 := OddNumbersGen(ctx, nums...)
	// выходной канал, где должны быть все числа из nums
	//outCh := FanIn(ctx, inputCh1, inputCh2)
	//for num := range outCh {
	//	fmt.Println(num) // 2, 0, 1, 4, 3, 6, 5 (у вас может быть другой порядок)
	//}

	nums := []Num{
		{0}, {1}, {2}, {3}, {4}, {5}, {6},
	}
	inputCh1 := EvenNumbersGen(ctx, nums...)
	inputCh2 := OddNumbersGen(ctx, nums...)
	// запись во временную очередь
	inCh := inTemp(ctx, inputCh1, inputCh2)
	// обработка из временной очереди и упорядочивание
	outCh := processTempCh(ctx, 2, inCh)
	
	for num := range outCh {
		fmt.Println(num) // {0} {1} {2} {3} {4} {5} {6} 
	}
}
