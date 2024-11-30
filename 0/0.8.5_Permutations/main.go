package main

import (
	"fmt"
	"sort"
)

func permute(prefix string, remaining string, result *[]string) {
	if len(remaining) == 0 {
		*result = append(*result, prefix)
		return
	}

	for i := 0; i < len(remaining); i++ {
		permute(prefix+string(remaining[i]), remaining[:i]+remaining[i+1:], result)
	}
}

func Permutations(input string) []string {

	var result []string
	permute("", input, &result)
	sort.Strings(result)
	return result
}

func main() {
	input := "abc"
	result := Permutations(input)
	fmt.Println(result)
}
