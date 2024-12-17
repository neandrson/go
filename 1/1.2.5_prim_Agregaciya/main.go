package main

import "fmt"

type Engine struct {
	Model string
}

type Car struct {
	Make   string
	Engine Engine
}

func main() {
	engine := Engine{Model: "V6"}
	car := Car{Make: "Toyota", Engine: engine}

	fmt.Printf("Car Make: %s\n", car.Make)
	fmt.Printf("Engine Model: %s\n", car.Engine.Model)
}
