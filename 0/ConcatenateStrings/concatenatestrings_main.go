package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(str1, str2 string) string {
	words := []string{str1, str2}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ConcatenateStrings("Ira", "Khomyakova"))
}
