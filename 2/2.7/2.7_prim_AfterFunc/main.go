package main

import (
	"fmt"
	"net/http"
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
/*func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel() // освобождает ресурсы, если slowOperation завершается до окончания тайм-аута
	return slowOperation(ctx), nil
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
}*/

// Прмер 3 - Паттерн для параллельных запросов с тайм-аутом
func fetchURL(url string, c chan string) {
	//создание http-клиента с тайм-аутом
	client := http.Client{
		/**
				  установка тайм-аута для HTTP-запроса; вы можете уменьшить или увеличить это время
				  - при уменьшении вы можете часто сталкиваться с тайм-аутом
		          - при увеличении данные будут получены за большее время
				  **/
		Timeout: 5 * time.Second,
	}
	//получение ответа по URL
	resp, err := client.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Ошибка при получении %s: %s", url, err)
		return
	}
	defer resp.Body.Close()
	c <- fmt.Sprintf("Ответ от %s: Статус - %s", url, resp.Status)
}

func main() {
	//создание массива URL-адресов
	urls := []string{
		"https://yandex.ru",
		"https://lyceum.yandex.ru",
		"https://translate.yandex.com",
		// Симуляция несуществующего URL
		"https://ihumaunkabir.com",
	}
	//создание канала C
	c := make(chan string, len(urls))
	//итерация по массиву URL-адресов
	for _, url := range urls {
		go fetchURL(url, c)
	}
	//установка общего тайм-аута для всех запросов
	timeout := time.After(15 * time.Second)
	//итерация до конца массива URL-адресов
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-c:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Произошел тайм-аут. Прерывание остальных запросов.")
			return
		}
	}
}
