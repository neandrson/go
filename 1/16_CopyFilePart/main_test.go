package main

import (
	"os"
	"testing"
)

func TestCopyFilePart(t *testing.T) {
	type test struct {
		fileContent    string
		fileName       string
		outFileName    string
		startPos       int
		wantError      bool
		outFileContent string
	}

	tests := []test{
		{
			fileContent:    `Go is a statically typed, compiled high-level programming language designed.`,
			fileName:       "file1.txt",
			outFileName:    "file1_out.txt",
			outFileContent: "is a statically typed, compiled high-level programming language designed.",
			startPos:       3},
		{
			fileContent:    `Golang outperforms Python in terms of microservices,APIs, and other fast-loading feature`,
			fileName:       "file2.txt",
			outFileName:    "file2_out.txt",
			outFileContent: `outperforms Python in terms of microservices, APIs, and other fast-loading feature`,
			startPos:       8},
		{
			fileContent: `Short test`,
			fileName:    "file1.txt",
			outFileName: "",
			startPos:    20,
			wantError:   true},
	}
	for _, tt := range tests {
		err := os.WriteFile
	}
}
