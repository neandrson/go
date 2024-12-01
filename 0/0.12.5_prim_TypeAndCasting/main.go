package main

import "fmt"

type Money float64

func (m Money) String() string {
	return fmt.Sprintf("$%.2f", m)
}

func main() {
	var salary Money = 2000.50

	fmt.Println(salary.String())
}
