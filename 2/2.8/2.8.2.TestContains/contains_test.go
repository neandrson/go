package main

import "testing"

func TestContains(t *testing.T) {

	testsCases := []struct {
		name    string
		numbers []int
		target  int
		want    bool
		err     error
	}{
		{
			name:    "Test1",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  5,
			want:    true,
			//err:  errors.New("tmp"),
		},
		{
			name:    "Test2",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target:  0,
			want:    false,
			//err:  errors.New("tmp"),
		},
	}

	for _, tc := range testsCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.numbers, tc.target)
			if got != tc.want {
				t.Errorf("Expected boolean:, %v, got: %v\n", tc.want, got)
			}

		})
	}
}
