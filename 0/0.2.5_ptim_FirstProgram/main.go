package main

import "fmt"

func main() {
	var age int
	fmt.Println("Введите ваш возраст:")
	fmt.Scanln(&age)
	fmt.Println(age)
}
