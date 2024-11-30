package main

import (
	students "_/C_/Users/neandr/go/0/10_students"
	"fmt"
	"lesson/students"
)

func (s Student) printData() {
	fmt.Printf("Name: %s, Age: %d\n", s.Name, s.Age)
}

func main() {
	student1 := students.Student{}
	user := students.User{}
	fmt.Println(student1)
}
