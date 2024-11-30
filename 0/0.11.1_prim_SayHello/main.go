package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Human interface {
	SayHello()
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	var h Human = Person{Name: "John", Age: 25}
	h.SayHello()
}
