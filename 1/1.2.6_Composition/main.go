package main

import "fmt"

type Engine struct {
	Model string
}

type Car struct {
	Make   string
	Engine // Встраивание структуры
}

func main() {
	car := Car{Make: "Ford", Engine: Engine{Model: "V8"}}

	fmt.Printf("Car Make: %s\n", car.Make)
	fmt.Printf("Engine Model: %s\n", car.Model) // Мы можем обращаться к Engine как к полю Car
}
