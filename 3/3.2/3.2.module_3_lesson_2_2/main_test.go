package main

import (
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		name      string
		input     []interface{}
		expected  []string
		stopAfter int
	}{
		{
			name:      "Test case 1",
			input:     []interface{}{1, 2, 3},
			expected:  []string{"1", "2", "3"},
			stopAfter: 0,
		},
		{
			name:      "Test case 2",
			input:     []interface{}{0, -1, 10},
			expected:  []string{"0", "-1", "10"},
			stopAfter: 0,
		},
		{
			name:      "Test case 3",
			input:     []interface{}{"", "str", "str2"},
			expected:  []string{"", "str"},
			stopAfter: 2,
		},
		{
			name:      "Test case 4",
			input:     []interface{}{0.3, -2.4, 0.0},
			expected:  []string{"0.3", "-2.4", "0"},
			stopAfter: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			done := make(chan struct{})
			if tc.stopAfter == 0 {
				defer close(done)
			}

			in := make(chan interface{})
			out := ToString(done, in)

			go func() {
				defer close(in)
				for i, num := range tc.input {

				}
			}()
		})
	}
}
