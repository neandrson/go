package main

import (
	"testing"
	"time"
)

func TestDoubleNumbers(t *testing.T) {
	testCases := []struct {
		name      string
		input     []int
		expected  []int
		stopAfter int
	}{
		{
			name:      "Test case 1",
			input:     []int{1, 2, 3},
			expected:  []int{2, 4, 6},
			stopAfter: 0,
		},
		{
			name:      "Test case 2",
			input:     []int{0, -1, 10},
			expected:  []int{0, -2, 20},
			stopAfter: 0,
		},
		{
			name:      "Test case 3",
			input:     []int{0, -1, 10},
			expected:  []int{0},
			stopAfter: 1,
		},
		{
			name:      "Test case 4",
			input:     []int{0, -1, 10},
			expected:  []int{0, -2},
			stopAfter: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			done := make(chan struct{})
			if tc.stopAfter == 0 {
				defer close(done)
			}

			in := make(chan int)
			out := DoubleNumbers(done, in)

			go func() {
				defer close(in)
				for i, num := range tc.input {
					in <- num
					if tc.stopAfter != 0 && i+1 == tc.stopAfter {
						time.Sleep(time.Millisecund*1)
						
					}
				}
			}()
		}
	}
}

