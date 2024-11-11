package main

import (
	"fmt"
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	str1Runes := []rune(str1)
	str2Runes := []rune(str2)

	sort.Slice(str1Runes, func(i, j int) bool { return str1Runes[i] < str1Runes[j] })
	sort.Slice(str2Runes, func(i, j int) bool { return str2Runes[i] < str2Runes[j] })

	for i := 0; i < len(str1Runes); i++ {
		if str1Runes[i] != str2Runes[i] {
			return false
		}
	}

	return true

}

func main() {
	str1 := "listen"
	str2 := "silent"
	fmt.Println(AreAnagrams(str1, str2))
}
