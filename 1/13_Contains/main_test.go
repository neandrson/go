package main

import (
	"fmt"
	"io"
)

type customReader struct {
}

func (cr *customReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func NewFakeReadOnlyFile(sizeBytes int) *FakeReadOnlyFile {
	return &FakeReadOnlyFile{bytes: sizeBytes}
}

// FakeReadOnlyFile simulates a fake read-only file of an arbitrary size.
// The contents of this file cycles through the lowercase alphabetical ascii characters (abcdefghijklmnopqrstuvwxyz)
// Data is generated as requested, so it is safe to generate a 100GB file with 1GB of memory free on your machine.
type FakeReadOnlyFile struct {
	bytes  int
	offset int
}

// Read reads lowercase ascii characters into b from f.
// n is number of bytes read.
// If a Read is attempted at the end of the file, io.EOF is returned.
func (f *FakeReadOnlyFile) Read(p []byte) (n int, err error) {
	if f.bytes == f.offset {
		return 0, io.EOF
	}
	bytesLeft := len(p)
	if (f.bytes - f.offset) < bytesLeft {
		bytesLeft = f.bytes - f.offset
	}
	return bytesLeft, err
}
