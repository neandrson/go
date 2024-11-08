package main

import (
	"fmt"
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (s Task) IsOverdue() bool {
	a := true
	if s.deadline < time.Now() {
		a = false
	}

	return a
}

func (s Task) IsTopPriority() bool {
	b := false
	if s.priority < 4 {
		b = true
	}
	return b
}

func main() {
	s := Task{summary: "abc", description: "abcdf", deadline: time.Now(), priority: 4}
	fmt.Println("дедлайн, %b", s.deadline)
	fmt.Println("приоритет %d", s.priority)
}
