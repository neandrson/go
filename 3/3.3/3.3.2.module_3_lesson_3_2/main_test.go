package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		fileContent []string
		files       []string
		expectError bool
		wordsCount  map[string]int
	}{
		{
			files: []string{"file1", "file2"},
			fileContent: []string{
				"Within hours a new European Super League",
				"Within days a new European Super league proposal had been released",
			},
			wordsCount: map[string]int{
				"within":   2,
				"hours":    1,
				"days":     1,
				"new":      2,
				"european": 2,
				"super":    2,
				"league":   2,
				"proposal": 1,
				"released": 1,
				"had":      1,
				"been":     1,
				"a":        2,
			},
			desc: "simple",
		},
		{
			files: []string{"file1", "file2"},
			fileContent: []string{
				"Some short text",
				"",
			},
			wordsCount: map[string]int{
				"some":  1,
				"short": 1,
				"text":  1,
			},
			desc: "short text",
		},
		{
			files: []string{"file1", "file2"},
			fileContent: []string{
				"",
				"",
			},
			wordsCount: map[string]int{},
		},
	}
}
