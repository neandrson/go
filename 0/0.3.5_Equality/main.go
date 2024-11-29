package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i string
	var j string
	var k string

	fmt.Scan(&i)
	fmt.Scan(&j)
	fmt.Scan(&k)

	l, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println("Некорректный ввод ")
		return
	}
	m, err := strconv.Atoi(j)
	if err != nil {
		fmt.Println("Некорректный ввод ")
		return
	}
	n, err := strconv.Atoi(k)
	if err != nil {
		fmt.Println("Некорректный ввод ")
		return
	}

	if l != m && m != n && n != l {
		fmt.Println("Все числа разные")
	} else if l == m && m == n && n == l {
		fmt.Println("Все числа равны")
	} else if l == m || m == n || n == l {
		fmt.Println("Два числа равны")
	}
}
