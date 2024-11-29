package main

import (
	"fmt"
)

func main() {
	var j int

	fmt.Scanln(&j)

	for i := j - 1; i > 0; i-- {
		j = j * i
		//fmt.Println(i, j)
	}
	fmt.Println(j)

}
