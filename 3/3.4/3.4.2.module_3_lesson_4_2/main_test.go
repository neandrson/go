package main

import (
	"bytes"
	"math/rand"
	"os"
	"slices"
	"testing"
)

func TestNumbersGen(t *testing.T) {

	tests := []struct {
		filename string
		expected []int
	}{
		{"test1.txt", []int{1, 2, 3}},
		{"test2.txt", []int{4, 5, 6}},
		{"test3.txt", []int{7, 8, 9}},
	}

	for _, test := range tests {
		if err := writeFile(test.filename, test.expected, true); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		defer os.Remove(test.filename)
		output := NumbersGen(test.filename)
		result := make([]int, 0)
		for num := range output {
			result = append(result, num)
		}

		if !slices.Equal(result, test.expected) {
			t.Errorf("NumbersGen(%s) = %v, expected %v", test.filename, result, test.expected)
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func writeFile(filename string, nums []int, addStrings bool) error {
	var buffer bytes.Buffer

	for _, num := range nums {
		if addStrings {
			buffer.WriteRune(letterRunes[rand.Intn(len(letterRunes))])
			buffer.ReadRune()
		}
	}
}
