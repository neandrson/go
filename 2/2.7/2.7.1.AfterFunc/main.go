package main

import (
	"context"
	"fmt"
	"time"
)

// Пример 1 - Использование тайм-аута во время выполнения функции
/*func main() {
//создание канала chan_c
chan_c := make(chan string)
//создание потока выполнения горутины
go func() {*/
/**
  создание фиктивной операции, которая занимает 3 секунды

  примечание: если вы хотите увидеть сообщение об успешном выполнении инструкций, сократите время до 1 секунды с миллисекундами или увеличьте время в параметре time.AfterFunc; также вместо Sleep() можно добавить операции, которые требуют больше времени
  **/
/*		time.Sleep(1 * time.Second)
		chan_c <- "Инструкции выполнены успешно."
	}()
	//создание тайм-аута, который не даёт функции выполняться дольше 2 секунд
	timeout := time.AfterFunc(2*time.Second, func() {
		chan_c <- "Время выполнения истекло."
	})

	result := <-chan_c
	fmt.Println(result)
	timeout.Stop() // Отмена функции тайм-аута
}*/

// Пример 2 - Контексты и отмена с тайм-аутом
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel() // освобождает ресурсы, если slowOperation завершается до окончания тайм-аута
	return slowOperation(ctx)
}

func main() {
	//создание контекста WithTimeout, который ограничивает продолжительность в течение 2 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	//создание канала chan_c
	chan_c := make(chan string)
	//создание потока выполнения горутины
	go func() {
		// создание фиктивной операции, которая занимает 3 секунды с использованием Sleep()
		time.Sleep(3 * time.Second)
		chan_c <- "Инструкции успешно завершены."
	}()

	select {
	case result := <-chan_c:
		fmt.Println(result)
	case <-ctx.Done():
		fmt.Println("Время операции истекло или было отменено.")
	}
}
