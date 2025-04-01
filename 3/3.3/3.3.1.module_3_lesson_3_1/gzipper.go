package gzipper

import (
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
				result := pattern.Find([]byte(filepath.Base(path)))
				if len(result) > 0 {
					jobs <- Work{path}
				}
			}
			return nil
		})
	}()
	return jobs
}

func compress(jobs <-chan Work) {
	for work := range jobs {
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
		err = os.Remove(work.Filepath)
		if err != nil {
			continue
		}
	}
}

func main() {
	dir := "C:\\"
	pattern := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`

	jobs := fs.ReadDir(dir, pattern)
	fmt.Println(regexp.MatchString(jobs))
}
