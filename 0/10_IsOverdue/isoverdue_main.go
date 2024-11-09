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
	return s.deadline.Compare(time.Now()) < 0
}

func (s Task) IsTopPriority() bool {
	b := false
	if s.priority >= 4 {
		b = true
	}
	return b
}

func main() {
	task := Task{summary: "Make Yandex Lyceum", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 10", priority: 3}
	fmt.Println(task.IsOverdue())
	fmt.Print(task.IsTopPriority())
}
