package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	//soundDog string
}

type Cat struct {
	//soundCat string
}

func (d Dog) MakeSound() string {
	fmt.Println("Гав!")
	return ""
}

func (c Cat) MakeSound() string {
	fmt.Println("Мяу!")
	return ""
}

func MakeSound(a Animal) {
	switch a.(type) {
	case Cat:
		fmt.Println("Мяу!")
	case Dog:
		fmt.Println("Гав!")
	default:
		fmt.Println("Unknown animal")
	}
}

func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}
