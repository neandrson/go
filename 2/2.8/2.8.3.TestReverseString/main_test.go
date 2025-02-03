package main

import "testing"

func TestReverseString(t *testing.T) {
	testsCases := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test1",
			str:  "abcd",
			want: "dcba",
		},
		{
			name: "Test2",
			str:  "efgh",
			want: "hgfe",
		},
	}

	for _, tc := range testsCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseString(tc.str)
			if got != tc.want {
				t.Errorf("Expected string:, %v, got: %v\n", tc.want, got)
			}

		})
	}
}
