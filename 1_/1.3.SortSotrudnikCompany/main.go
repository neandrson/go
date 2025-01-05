package main

import (
	"fmt"
)

type CompanyInterface interface {
	AddWorkerInfo(Name, Position string, Salary, Experience uint) error
	SortWorkers() ([]string, error)
}

type AddWorkerInfo struct {
	Name, Position     []string
	Salary, Experience []uint
}

func SortWorkers(Salary []uint, Position []string) []string {
	return "" //sort.Sort(string(Salary)) + sort.Sort(Position)
}

func (a AddWorkerInfo) AddWorkerInfo(Name, Position []string, Salary, Experience []uint) {
	fmt.Println(Name, Position, Salary, Experience)
}

func main() {
	workers := []AddWorkerInfo{
		{Name: "Михаил", Position: "директор", Salary: 200, Experience: 5},
		{Name: "Игорь", Position: "зам. директора", Salary: 180, Experience: 3},
		{Name: "Николай", Position: "начальник цеха", Salary: 120, Experience: 2},
		{Name: "Андрей", Position: "мастер", Salary: 90, Experience: 10},
		{Name: "Виктор", Position: "рабочий", Salary: 80, Experience: 3},
	}
	workersSort, err := AddWorkerInfo(workers)
	if err != nil {
		fmt.Println("")
	}
	fmt.Println(workersSort)
}
