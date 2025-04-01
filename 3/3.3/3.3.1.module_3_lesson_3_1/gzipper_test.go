package gzipper

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestCompress(t *testing.T) {
	dir := "./"

	testFiles := []struct {
		desc        string
		fileContent string
		fileName    string
	}{
		{
			desc:        "",
			fileContent: "Test file 1",
			fileName:    "file1.txt",
		},
		{
			desc:        "",
			fileContent: "Test file 2",
			fileName:    "file2.txt",
		},
	}

	for _, testFile := range testFiles {
		path := filepath.Join(dir, testFile.fileName)
		createTestFile(t, path, testFile.fileContent)
	}

	defer func() {
		for _, testFile := range testFiles {
			path := filepath.Join(dir, testFile.fileName)
			pathGzip := filepath.Join(dir, testFile.fileName+".gz")
			os.Remove(path)
			os.Remove(pathGzip)
		}
	}()

	// Define the file pattern to match
	pattern := regexp.MustCompile("\\.txt$")

	// Generate a stream of files to compress
	files := FileNameGen(dir, pattern)

	// Compress the files in parallel
	compress(files)

	// Verify that the compressed
}

func createTestFile(t *testing.T, path string, fileContent string) {

}
