package gzipper

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

type Work struct {
	Filepath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work)
	go func() {
		defer close(jobs)
		filepath.Walk(dir, func(path string, finfo fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !finfo.IsDir() {
				result := pattern.Find(
					[]byte(filepath.Base(path)),
				)
				if len(result) > 0 {
					jobs <- Work{path}
				}
			}
			return nil
		},
		)
	}()
	return jobs
}

func compress(jobs <-chan Work) {
	for work := range jobs {
		if work.Filepath == "input.txt" || work.Filepath == "output.txt" {
			continue
		}
		out, err := os.Create(work.Filepath + ".gz")
		if err != nil {
			continue
		}
		defer out.Close()

		gw := gzip.NewWriter(out)
		defer gw.Close()

		in, err := os.Open(work.Filepath)
		if err != nil {
			continue
		}
		defer in.Close()

		_, err = io.Copy(gw, in)
		if err != nil {
			continue
		}
	}
}

func createTestFile(path string, fileContent string) {
	//limitReader := io.LimitReader(*testFile.Reader, testFile.FileSize)

	f, err := os.Create(path)
	if err != nil {
		fmt.Errorf("file creation error: %v", err)
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
			fmt.Errorf("writing temp file: %+v", err)
		}
	}

	if err := w.Flush(); err != nil {
		t.Errorf("flushing file: %+v", err)
	}

	//return
}

func main() {
	dir := "./"

	files := []struct {
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

	for _, file := range files {
		path := filepath.Join(dir, file.fileName)
		createTestFile(path, file.fileContent)
	}

	defer func() {
		for _, file := range files {
			path := filepath.Join(dir, file.fileName)
			pathGzip := filepath.Join(dir, file.fileName+".gz")
			os.Remove(path)
			os.Remove(pathGzip)
		}
	}()

	// Define the file pattern to match
	pattern := regexp.MustCompile("\\.txt$")

	// Generate a stream of files to compress
	fil := FileNameGen(dir, pattern)

	// Compress the files in parallel
	compress(files)

	jobs := fs.ReadDir(dir, pattern)
	fmt.Println(regexp.MatchString(jobs))
}
