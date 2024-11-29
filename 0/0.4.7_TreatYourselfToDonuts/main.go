package main

import (
	"fmt"
)

func main() {
	var j int
	var k int

	fmt.Scan(&j)

	//l, err := strconv.Atoi(j)
	//if j < 0 {
	//	fmt.Println("Некорректный ввод ")
	//	return
	//}
	for i := 1; i <= j; i++ {

		if i%3 == 0 || i%5 == 0 {
			//fmt.Println(i)
			continue
		}
		//j := 0
		k = k + i
		//fmt.Println(k)
	}
	fmt.Println(k)
}
