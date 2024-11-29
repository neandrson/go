package main

import (
	"fmt"
)

func main() {
	var j int
	var k int

	fmt.Scanln(&j)

	//l, err := strconv.Atoi(j)
	if j < 0 {
		fmt.Println("Некорректный ввод ")
		return
	}
	for i := j; i > 0; i-- {

		if i%2 == 0 {
			continue
		}
		//j := 0
		k = k + i
		//fmt.Println(i, k)
	}
	fmt.Println(k)
}
