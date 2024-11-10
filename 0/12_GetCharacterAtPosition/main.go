package main

import (
	"errors"
	"fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	fmt.Println(len(str))
	fmt.Println(position)
	a := []rune(str)
	fmt.Println(position)
	if position >= len(str) {
		return 0, errors.New("position out of range")
	}
	return a[position], nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 14))
}
