package main

import (
	"testing"
)

func TestMultiplyPipeline(t *testing.T) {
	testCases := []struct {
		inputNums [][]int
		expected  int
	}{
		{
			inputNums: [][]int{{1, 2, 3}, {4, 5, 6}},
			expected:  720,
		},
		{
			inputNums: [][]int{{-1, -2, -3}, {4, 5, 6}},
			expected:  120,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		actual := MultiplyPipeline(tc.inputNums...)
		if actual != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, actual)
		}
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, -2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
	}

	for _, tc := range testCases {
		in := make(chan int)
		out := Filter(in)

		go func() {
			for _, num := range tc.input {
				in <- num
			}
			close(in)
		}()

		actual := []int{}
		for num := range out {
			actual = append(actual, num)
		}
	}
}
