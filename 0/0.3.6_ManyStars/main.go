package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i string
	//var j string
	//var k string

	fmt.Scan(&i)
	//fmt.Scan(&j)
	//fmt.Scan(&k)

	l, err := strconv.Atoi(i)
	if err != nil || l < 0 {
		fmt.Println("Некорректный ввод ")
		return
	}
	/*m, err := strconv.Atoi(j)
	if err != nil || m < 0 {
		fmt.Println("Некорректный ввод ")
		return
	}
	n, err := strconv.Atoi(k)
	if err != nil || n < 0 {
		fmt.Println("Некорректный ввод ")
		return
	}*/

	if l < 10 {
		fmt.Println("Число меньше 10")
	} else if l >= 10 && l < 100 {
		fmt.Println("Число меньше 100")
	} else if l >= 100 && l < 1000 {
		fmt.Println("Число меньше 1000")
	} else if l >= 1000 {
		fmt.Println("Число больше или равно 1000")
	}
}
