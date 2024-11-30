package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(input string) bool {
	input = strings.ToLower(input) // Нормализуем строку до нижнего регистра
	input = strings.Replace(input, " ", "", -1)
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(" a b c       cba  "))
}
