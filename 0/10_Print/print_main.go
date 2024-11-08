package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (s Person) Print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s", s.name, s.age, s.address)
}

func main() {
	man := Person{name: "Anton", age: 31, address: "Krasnogorsk"}
	man.Print()
}
