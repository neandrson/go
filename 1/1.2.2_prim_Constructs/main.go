package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Конструктор в Go обычно возвращает новый экземпляр структуры
func NewPerson(name string) Person {
	return Person{Name: name, Age: 40}
}

func main() {
	// Используем конструктор
	person := NewPerson("Alice")

	fmt.Printf("Имя: %s, Возраст: %d\n", person.Name, person.Age)
}
