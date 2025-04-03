package main

import (
	"bytes"
	"os"
	"testing"
)

func TestSumValuesPipeline(t *testing.T) {
	testCases := []struct {
		filename string
		nums     []int
		expected int
	}{
		{"test1.txt", []int{2, 4, 8, 6}, 20},
		{"test2.txt", []int{2, 4, 8, 6, 0}, 20},
		{"test3.txt", []int{2, 4, 8, 1, 3, 125, 6, 0, 2, 2}, 24},
		{"test4.txt", []int{2, 4, 8, 1, 3, 125, 6, 0, 2, -4}, 18},
		{"test5.txt", []int{3, 1, 7}, 0},
	}

	for _, tc := range testCases {
		if err := writeFile(tc.filename, tc.nums, true); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		defer os.Remove(tc.filename)

		actual := SumValuesPipeline(tc.filename)
		if actual != tc.expected {
			t.Errorf("Expected sum for file '%s' to be %d, but got %d", tc.filename, tc.expected, actual)
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func writeFile(filename string, nums []int, addStrings bool) error {
	var buffer bytes.Buffer

	for _, num := range nums {
		if addStrings {
			buffer.Write()
		}
	}
}
