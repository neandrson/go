package main

import (
	"bufio"
	"os"
	"strconv"
)

func NumbersGen(filename string) <-chan int {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}

	out := make(chan int)
	go func() {
		defer f.Close()
		defer close(out)
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			num, err := strconv.Atoi(fs.Text())
			if err != nil {
				continue
			}
			out <- num
		}
	}()

	return out
}

func main() {

}
