package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string, age int) *User {

	return &User{
		Name:   name,
		Age:    age,
		Active: true,
	}
}

func main() {
	user1 := NewUser("Alice", 25)
	user2 := NewUser("Bob", 30)

	fmt.Println(user1)
	fmt.Println(user2)
}
