package main

import "fmt"

type Logger interface {
	Log()
}

type LogLevel string

const Error LogLevel = "Error"
const Info LogLevel = "Info"

type Log struct {
	Level LogLevel
}

func (l Log) Log(s string) {
	if l.Level == Error {
		fmt.Printf("ERROR: %s", s)
	} else {
		fmt.Printf("INFO: %s", s)
	}
}

func main() {
	errorLog := &Log{Level: Error}
	errorLog.Log("This is an error message")
}
