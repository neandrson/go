package main

import (
	"fmt"
	"unicode"
)

func DeleteVowels(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		switch unicode.ToLower(rune(s[i])) {
		case 'a':
			continue
		case 'e':
			continue
		case 'i':
			continue
		case 'o':
			continue
		case 'u':
			continue
		}
		result += string(s[i])
	}
	return result
}

func main() {
	s := "aertgwer'gjmfokvbniopu"
	fmt.Println(DeleteVowels(s))
}
