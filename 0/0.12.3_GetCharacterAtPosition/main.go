package main

import (
	"errors"
	"fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	a := []rune(str)
	if position < 0 || position >= len(a) {
		return 0, errors.New("position out of range")
	}
	return []rune(str)[position], nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 14))
}
