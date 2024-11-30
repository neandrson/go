package main

import "fmt"

func main() {
	/*var age int
	fmt.Println("Введите ваш возраст:")
	fmt.Scanln(&age)
	fmt.Println(age)*/
	var age int
	fmt.Println("Введите ваш возраст:")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(age)
}

/*package main // создаём пакет main для нашей программы

import "fmt" // подключаем библиотеку fmt для работы со строками и консолью

func main() { // обозначаем точку входа в нашу программу
	var age int // определяем переменную для хранения возраста
	fmt.Println("Введите ваш возраст:") // выводим сообщение с просьбой ввести возраст
	fmt.Scanln(&age)  // просим пользователя ввести число в консоль
	fmt.Println(age)  // выводим в консоль его же возраст
}*/
