package main

import (
	"fmt"
	"time"
)

// тип «таймер» представляет одно событие
// когда время до окончания работы таймера истекает, текущее время будет отправлено в канал C, если только таймер не был создан с помощью AfterFunc
// таймер должен быть создан с помощью NewTimer или AfterFunc
/*type Timer struct {
	C <-chan Time
	r runtimeTimer
}*/

// Пример 2
func timersStoppage(v time.Timer) {
	<-v.C
	fmt.Println("Второй таймер сработал")
}

// Пример 1
/*func main() {
	// создание таймера, время которого истечёт через 3 секунды
	timer := time.NewTimer(3 * time.Second)
	// канал C отправляет значение, которое указывает на окончание работы таймера
	<-timer.C
	fmt.Println("Time's up!")
}*/

// Пример 2
/*func main() {
	// создание первого таймера
	timer := time.NewTimer(3 * time.Second)
	// в канал C отправляется значение, которое указывает на окончание работы таймера
	<-timer.C
	fmt.Println("Первый таймер сработал!")
	// создание второго таймера
	timer_s := time.NewTimer(time.Second)
	// создание горутины
	go timersStoppage(*timer_s)
	// удалите // из строки ниже, чтобы сработал второй таймер
	// timersStoppage((*timer_s))
	// остановка второго таймера перед срабатыванием
	stop_s := timer_s.Stop()
	if stop_s {
		fmt.Println("Второй таймер остановлен")
	}
}*/

// Пример 3
func timeoutTest() string {
	// уменьшаем время здесь, чтобы попасть в пределы
	time.Sleep(2 * time.Second)
	return "Функция TimeoutTest выполнена!"
}

func main() {
	// создание канала C
	c := make(chan string, 1)
	// создание потока выполнения горутины
	go func() {
		str := timeoutTest()
		c <- str
	}()
	// создание тайм-аута для выполнения функции
	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Время вышло!")
	}
}
