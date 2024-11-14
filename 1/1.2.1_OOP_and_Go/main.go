package main

import "fmt"

type Cat struct {
	Name string
}

// Конструктор. Функция, которая создаёт экземпляр нашего объекта
func NewCat(name string) Cat {
	return Cat{Name: name}
}

// Метод Meow() — то есть «мяукать»
func (c Cat) Meow() {
	fmt.Printf("%s мяукнул(а)!\n", c.Name)
}

func main() {
	myCat := NewCat("Барсик")
	myCat.Meow()
}
