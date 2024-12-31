package main

import (
	"slices"
	"testing"
)

func TestSortNames(t *testing.T) {
	tests := []struct {
		names    []string
		expected []string
	}{
		{
			names:    []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"},
			expected: []string{"Аксинья", "Арина", "Валентина", "Варвара", "Влада", "Жанна", "Ярослава"},
		},
		{
			names:    []string{"Мария", "Фрося", "Раиса"},
			expected: []string{"Мария", "Раиса", "Фрося"},
		},
	}
	for _, tc := range tests {
		SortNames(tc.names)
		if slices.Compare(tc.expected, tc.names) != 0 {
			t.Errorf("TestSortNames failed. Expected: %v, Got: %v", tc.expected, tc.names)
		}
	}
}
