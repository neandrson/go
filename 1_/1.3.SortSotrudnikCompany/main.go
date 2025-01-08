package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CompanyInterface interface {
	AddWorkerInfo(Name, Position string, Salary, Experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

func AddWorkerInfo(workers Worker) error {
	var worker []string
	var sum uint = 0
	sum = workers.Salary * workers.Experience * 12
	worker = append(worker, workers.Name+" - "+strconv.FormatUint(uint64(sum), 10)+" - "+workers.Position)
	fmt.Println(strings.Trim(fmt.Sprintf("%v", worker), "[]"))
	return nil
}

func SortWorkers(workers []string) ([]string, error) {
	return workers, nil //sort.Sort(string(Salary)) + sort.Sort(Position)
}

func main() {
	workers := []Worker{
		{Name: "Михаил", Position: "директор", Salary: 200, Experience: 5},
		{Name: "Игорь", Position: "зам. директора", Salary: 180, Experience: 3},
		{Name: "Николай", Position: "начальник цеха", Salary: 120, Experience: 2},
		{Name: "Андрей", Position: "мастер", Salary: 90, Experience: 10},
		{Name: "Виктор", Position: "рабочий", Salary: 80, Experience: 3},
	}
	for _, worker := range workers {
		err := AddWorkerInfo(worker)
		if err != nil {
			fmt.Println("Ошибка")
		}
	}
}
