package main

import (
	"testing"
)

func TestTee(t *testing.T) {
	tests := []struct {
		name            string
		input           []int
		expectedOutput1 []int
		expectedOutput2 []int
	}{
		{
			name:            "Empty input",
			input:           []int{},
			expectedOutput1: nil,
			expectedOutput2: nil,
		},
		{
			name:            "Non-empty input",
			input:           []int{1, 2, 3, 4, 5},
			expectedOutput1: []int{1, 2, 3, 4, 5},
			expectedOutput2: []int{1, 2, 3, 4, 5},
		},

		{
			name:            "Non-empty input 2",
			input:           []int{5, 5, 1, 7, 9, 2},
			expectedOutput1: []int{5, 5, 1, 7, 9, 2},
			expectedOutput2: []int{5, 5, 1, 7, 9, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			done := make(chan struct{})
			defer close(done)

			in := make(chan int)
			out1, out2 := Tee(done, in)

			go func() {
				for _, data := range test.input {
					in <- data
				}
				close(in)
			}()

			var result1 []int
			var result2 []int

			for val1 := range in {
				result1 = append(result1, val1)
			}
			for val2 := range in {
				result2 = append(result2, val2)
			}
			if out1 != in {
				t.Errorf("Expacted: %d, do got: %d", test.expectedOutput1, out1)
			}
			if out2 != in {
				t.Errorf("Expacted: %d, do got: %d", test.expectedOutput2, out2)
			}
		})
	}
}
