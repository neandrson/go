package main

import (
	"fmt"
)

func IsPowerOfTwoRecursive(N int) {
	if N%2 == 0 || N == 1 {
		//fmt.Println(N)
		if N/2 == 1 || N == 1 {
			fmt.Println("YES")
			return
		}
		IsPowerOfTwoRecursive(N / 2)
		//return
	}
	if N%2 != 0 {
		//fmt.Println(N)
		fmt.Println("NO")
		//return
	}
}

func main() {
	IsPowerOfTwoRecursive(3)
}
