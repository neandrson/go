package main

import "fmt"

type Employee struct {
	name     (string)
	position (string)
	salary   (float64)
	bonus    (float64)
}

func (s Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\nPosition: %s\nTotal Salary: %0.2F", s.name, s.position, s.salary+s.bonus)
}

func main() {
	man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
	man.CalculateTotalSalary()
}
