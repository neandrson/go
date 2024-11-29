package main

import (
	"fmt"
	//"strconv"
)

func main() {
	var i int

	fmt.Scan(&i)

	for j := 1; j <= 10; j++ {
		fmt.Println(i, "*", j, "=", i*j)
	}

}
