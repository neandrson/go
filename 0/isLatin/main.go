package main

import "fmt"

func isEnglishLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// isLatin проверяет, состоят ли все символы в строке input из букв английского алфавита [1](https://otvet.mail.ru/question/240299301)
func isLatin(input string) bool {
	for _, r := range input {
		if !isEnglishLetter(r) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isLatin("Andrey"))
}
