package main

import "fmt"

func merge(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func Permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	runes := []rune(input)
	subPermutations := Permutations(string(runes[0 : len(input)-1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
		fmt.Println(result)
	}

	return result
}

func PermutationsTailRecursive(input string) []string {
	var helper func(input []rune, previousResult []string) []string
	helper = func(input []rune, previousResult []string) []string {
		if len(input) == 0 {
			return previousResult
		} else {
			newResult := []string{}
			for _, element := range previousResult {
				newResult = append(newResult, merge([]rune(element), input[0])...)
			}
			return helper(input[1:], newResult)
		}
	}

	runes := []rune(input)
	return helper(runes[1:], []string{string(runes[0])})
}

func main() {
	arr := "abc"
	fmt.Println(Permutations(arr))
}
