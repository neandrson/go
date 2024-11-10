package main

import "fmt"

func ConcatStringsAndInt(str1, str2 string, num int) string {
	return str1 + " " + str2 + " " + string(num)
}

func main() {
	fmt.Println(ConcatStringsAndInt("al", "himik", 239))
}
