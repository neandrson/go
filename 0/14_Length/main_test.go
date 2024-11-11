package main

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{
	{-1, "negative"},
	{5, "short"},
}

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}
