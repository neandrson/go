package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	testsCases := []struct {
		name   string
		str    string
		target string
		want   bool
	}{
		{
			name:   "Test1",
			str:    "abcd",
			target: "abcd",
			want:   true,
		},
		{
			name:   "Test2",
			str:    "efgh",
			target: "abcd",
			want:   false,
		},
		{
			name:   "Test3",
			str:    "",
			target: "abcd",
			want:   false,
		},
		{
			name:   "Test4",
			str:    "",
			target: "",
			want:   true,
		},
	}

	for _, tc := range testsCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := AreAnagrams(tc.str, tc.target)
			if got != tc.want {
				t.Errorf("Expected string:, %v, got: %v\n", tc.want, got)
			}
		})
	}
}
