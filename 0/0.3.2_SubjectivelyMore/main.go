package main

import (
	"fmt"
)

func main() {
	var i int
	var j int
	//fmt.Print("Введите Первое и Второе числа ")
	fmt.Scan(&i)
	fmt.Scan(&j)
	if i > j {
		fmt.Println("Первое число больше второго")
	} else if i < j {
		fmt.Println("Второе число больше первого")
	} else if i == j {
		fmt.Println("Числа равны")
	}
}
