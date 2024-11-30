package main

import (
	"fmt"
	"time"
)

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (s Task) IsOverdue() bool {
	return time.Now().After(s.deadline)
}

func (s Task) IsTopPriority() bool {
	b := false
	if s.priority >= 4 {
		b = true
	}
	return b
}

func (s *ToDoList) TasksCount() int {
	return len(s.tasks)
}

func (s *ToDoList) NotesCount() int {
	return len(s.notes)
}

func (s *ToDoList) CountTopPrioritiesTasks() int {
	count := 0
	for _, task := range s.tasks {
		if task.priority >= 4 {
			count++
		}
	}
	return count
}

func (s *ToDoList) CountOverdueTasks() int {
	count := 0
	for _, task := range s.tasks {
		if task.IsOverdue() {
			count++
		}
	}
	return count
}

func main() {
	todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}, Task{summary: "Make Yandex Lyceum Task 10", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 3}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
	fmt.Println(todo.TasksCount())
	fmt.Println(todo.NotesCount())
	fmt.Println(todo.CountTopPrioritiesTasks())
	fmt.Print(todo.CountOverdueTasks())
}
