package main

import "testing"

type Test struct {
	input    []byte
	expected int
	err      error
}

func TestGetUTFLength(t *testing.T) {
	tests := []Test{
		// Тест с корректной строкой UTF-8.
		{input: []byte("Hello, 世界"), expected: 9, err: nil},
		// Тест с некорректной строкой UTF-8.
		{input: []byte{0x80}, expected: 0, err: ErrInvalidUTF8},
		// Тест с пустой строкой.
		{input: []byte{}, expected: 0, err: nil},
		// Тест с одной руной UTF-8.
		{input: []byte("你"), expected: 1, err: nil},
	}
	for _, tt := range tests {
		got, err := GetUTFLength(tt.input)
		if got != tt.expected {
			t.Errorf("GetUTFLength() got = %v, expected = %v", got, tt.expected)
		}
		if err != tt.err {
			t.Errorf("GetUTFLength() error = %v, expected = %v", err, tt.err)
		}
	}
}
