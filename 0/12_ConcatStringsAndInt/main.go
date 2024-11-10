package main

import (
	"fmt"
	"strconv"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + str2 + strconv.Itoa(num)
}

func main() {
	fmt.Println(ConcatStringsAndInt("al", "himik", 239))
}
