package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "АнДРей"
	lower := strings.ToLower(name)
	fmt.Println(lower) // Output: андрей

	words := []string{"Это", "несколько", "слов"}
	joined := strings.Join(words, " ")
	fmt.Println(joined) // Output: Это несколько слов

	text := "Я ЛЮБЛЮ Golang"
	fmt.Println(strings.Replace(text, "Golang", "GO", -1))
	// Output: Я ЛЮБЛЮ GO
}
