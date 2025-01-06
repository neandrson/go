package main

import (
	"fmt"
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

func (w Worker) AddWorkerInfo(workers []Worker) error {
	//var workerSlice []string
	/*for _, result := range worker {
		workerSlice = append(workerSlice, string(result))
	}*/
	fmt.Println(w.Name, w.Position, w.Salary, w.Experience)
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
	worker, _ := AddWorkerInfo(workers)
	/*if err != nil {
		fmt.Println("Ошибка")
	}*/
	fmt.Println(worker)
}
