package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// создадим тип данных для хранения id пользоваталя
type userID string

func readSource(ctx context.Context) error {
	// имитируем долгую работу функции
	time.Sleep(3 * time.Second)
	// допустим, возникла ошибка в процессе
	return fmt.Errorf("some error in readSource\n")
}

func processSourceData(ctx context.Context) error {
	// получаем данные в цикле
	for {
		select {
		// раз в секунду получаем новые данные
		case <-time.After(time.Second):
			// здесь может быть код получения очередной порции данных
			fmt.Println("process data bit by bit...")
		// проверим контекст на отмену
		case <-ctx.Done():
			fmt.Println("processSourceData was canceled")
			return nil
		}
	}
}

/*func ProcessRequest(userID string) {
	// Пример 1
	// сохраним значение в контексте
	ctx := context.WithValue(context.Background(), "userID", userID)
	// функция обработки
	HandleResponse(ctx)

	// Пример 2
	// сохраним значение myValue по ключу myKey
	ctx = context.WithValue(ctx, "myKey", "myValue")
	fmt.Printf("my value is %s\n", ctx.Value("myKey")) // my value is myValue
	// сохраним значение anotherValue по ключу myKey
	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	fmt.Printf("my value is %s\n", anotherCtx.Value("myKey")) // my value is anotherValue
}*/

// Пример 3
func ProcessRequest(id userID) {
	// сохраним значение в контексте
	ctx := context.WithValue(context.Background(), "userID", id)
	id, ok := ctx.Value("userID").(userID)
	if !ok {
		// другой тип объекта
	}
}

// здесь контекст уже содержит userID
func HandleResponse(ctx context.Context) { // допустим, переданный контекст содержит значение по ключу userID
	// Пример 1
	//fmt.Printf("handling response for (%v)", ctx.Value("userID"))

	//Пример 2
	newUserID := "22"
	ctx = context.WithValue(ctx, "userID", newUserID)             // запишем значение по ключу userID
	fmt.Printf("handling response for (%v)", ctx.Value("userID")) // handling response for 22
}

func main() {
	// Пример 1 Отмена контекста
	/*ctx := context.Background()
	// ожидаем завершения горутин
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// запускаем функцию обработки данных
		if err := processSourceData(ctx); err != nil {
			fmt.Printf("processSourceData(ctx): %s", err)
		}
	}()
	go func() {
		defer wg.Done()
		// запускаем функцию чтения данных
		if err := readSource(ctx); err != nil {
			fmt.Printf("readSource(ctx): %s", err)
		}
	}()
	// ждём завершения
	wg.Wait()*/

	// Пример 2
	// контекст, который можно отменить
	ctx := context.Background()
	ctxWithCancel, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := processSourceData(ctxWithCancel); err != nil {
			fmt.Printf("processSourceData(ctxWithCancel): %s", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := readSource(ctxWithCancel); err != nil {
			// при ошибке в функции чтения подадим сигнал через контекст
			cancelCtx()
			fmt.Printf("readSource(ctxWithCancel): %s", err)
		}
	}()
	wg.Wait()
}
