package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

/*func ProcessData(ctx context.Context, bucket, fileName string) error{
	// Клиент для работы с хранилищем
	client, err := storage.NewClient(ctx)
	if err != nil {
		// TODO: handle error.
	}
	// Откроем файл в Google Cloud Storage
	file, err := client.Bucket(bucket).Object(filename).NewReader(ctx)
	if err != nil {
		// TODO: handle error.
	}
	// Здесь обрабатываем данные
	...
}*/

func OpenLocalFile(filename string) (*os.File, error) {
	// Открываем локальный файл
}
func OpenGCSFile(
	ctx context.Context,
	bucket, fileName string,
) (*storage.Reader, error) {
	// Открываем файл из Google Cloud Storage
}

func ProcessData(ctx context.Context, reader io.Reader) error {
	data := make([]byte, 1024)          // Создадим буфер для чтения данных в него
	bytesRead, err := reader.Read(data) // Прочитаем данные в буфер
	if err != nil {
		// TODO: handle error.
	}
	// Сколько прочитали байт и сам контент
	fmt.Printf("Прочитано %d байт: %s", bytesRead, string(data[:bytesRead]))
	return nil
}
