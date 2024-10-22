package main

import (
	"fmt"
)

func main() {
	var k int

	fmt.Scan(&k)

	if k <= 1 {
		return
	}

	var n2, n1 = 0, 1

	for i := 1; i <= 10; i++ {
		if n1+n1+n2 >= k && n1 >= k {
			fmt.Println(n1)
		} else if i >= 1 {
			i--
		}
		n2, n1 = n1, n1+n2

	}
}
