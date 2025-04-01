package gzipper

import (
	"bufio"
	"fmt"
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
	//limitReader := io.LimitReader(*testFile.Reader, testFile.FileSize)

	f, err := os.Create(path)
	if err != nil {
		t.Errorf("file creation error: %v", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("f.Close() error:%+v\n", err)
		}
	}()

	w := bufio.NewWriter(f)
	buf := make([]byte, len(fileContent)-1)
	for {
		/*nBuf, err := limitReader.Read(buf)
		if err == io.EOF {
			break
		}*/
		// fmt.Printf("nBuf: %d, err: %+v, data: %s\n", nBuf, err, base64.StdEncoding.EncodeToString(buf))

		if _, err = w.Write(buf[:len(fileContent)-1]); err != nil {
			t.Errorf("writing temp file: %+v", err)
		}
	}

	if err := w.Flush(); err != nil {
		t.Errorf("flushing file: %+v", err)
	}

	//return
}
