package main

import (
	"testing"
)

type Test struct {
	old string
	new string
}

var tests = []Test{
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"},
	{"zyxwvutsrqponmlkjihgfedcbaZYXWVUTSRQPONMLKJIHGFEDCBA", "zyxwvtsrqpnmlkjhgfdcbZYXWVTSRQPNMLKJHGFDCB"},
	{"Hello, world!", "Hll, wrld!"},
	{"1234567890ao", "1234567890"},
}

func TestDeleteVowels(t *testing.T) {
	for i, test := range tests {
		line := DeleteVowels(test.old)
		if line != test.new {
			t.Errorf("#%d: DeleteVowels(%s)=%s; want %s", i, test.old, line, test.new)
		}
	}
}
