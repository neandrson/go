package main

import "fmt"

var subPermutations []string

func merge(ins []rune, c rune) (result []string) {
	for i := len(ins); i >= 0; i-- {
		//fmt.Println(result)
		//fmt.Println(i)
		//fmt.Println(ins[:i])
		//fmt.Println(c)
		fmt.Println(ins[i:])
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
		//fmt.Println(result)
	}
	return
}

func Permutations(input string) []string {

	if len(input) <= 1 {
		return []string{input}
	}
	//fmt.Println(input)
	runes := []rune(input)
	//fmt.Println(runes)

	for i := 0; i < len(input)-1; i++ {
		subPermutations = Permutations(string(runes[i : len(input)-1]))
		//fmt.Println(subPermutations)
	}

	result := []string{}
	for _, s := range subPermutations {
		//fmt.Println(merge([]rune(s), runes[len(input)-1]))
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
		//fmt.Println(result)
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

/*func Perm(input string) {
	a := []rune(input)
	f (a)
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}*/

/*func join(ins []rune, c rune) (result []string) {
	for i := len(ins); i >= 0; i-- {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:])) //
	}
	return
}

func Permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}*/

func main() {
	arr := "abc"
	fmt.Println(Permutations(arr))
	/*Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})*/

}
