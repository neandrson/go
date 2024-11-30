package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Scanln(&score)
	if score >= 60 {
		fmt.Println("Отличная работа!")
	} else {
		fmt.Println("Нужно усерднее учиться!")
	} // Выводит "Отличная работа!"
}
