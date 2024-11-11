package main

import "testing"

type Test struct {
	a int
	b int
}

var tests = []Test{
	{-1, 5},
	{10, 20},
	{0, 5},
	{90, 0},
	{101, 0},
}

func TestMultiply(t *testing.T) {
	for _, test := range tests {
		size := Multiply(test.a, test.b)
		if size == 0 {
			t.Errorf("Error not nule %s", test.b)
		}
	}
}
