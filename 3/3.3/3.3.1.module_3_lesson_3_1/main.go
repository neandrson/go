package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

type Work struct {
	FilePath string
}

func FileNameGen(dir string, pattern *regexp.Regexp) <-chan Work {
	jobs := make(chan Work)
	m, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Printf("File in path no %v", err)
		return nil
	}

	for _, val := range m {
		go func() {
			defer close(jobs)
			/*e := filepath.Walk(dir, func(path string, d fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				yandexIssue := regexp.MustCompile("^.+\\.(txt)$")
				if !yandexIssue.Match([]byte(path)) && !d.IsDir() && pattern.MatchString(path) {
					jobs <- Work{FilePath: path}
				}
				return nil
			})
			if e != nil {
				log.Fatal(e)
			}*/

			fmt.Println(val)

		}()
	}
	return jobs
}

func compress(jobs <-chan Work) {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range jobs {
				file, err := os.Open(w.FilePath)
				if err != nil {
					panic(err)
				}

				fw, err := os.Create(w.FilePath + ".gz")
				if err != nil {
					panic(err)
				}

				cw := gzip.NewWriter(fw)
				_, err = io.Copy(cw, file)

				if err != nil {
					panic(err)
				}
				fw.Close()
				file.Close()
				cw.Close()
			}
		}()
	}
	wg.Wait()
}

func main() {
	dir := "C:\\"
	pattern := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`

	jobs := fs.ReadDir(dir, pattern)
	fmt.Println(regexp.MatchString(jobs))
}
