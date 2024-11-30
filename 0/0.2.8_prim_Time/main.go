package main

import (
	"fmt"
	"time" // импортируем библиотеку time
)

func main() {
	now := time.Now() // вызываем команду, которая кладёт в переменную текущее время
	fmt.Println(now)
}
