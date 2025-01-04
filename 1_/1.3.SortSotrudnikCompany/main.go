package main

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type AddWorkerInfo struct {
	name, position     string
	salary, experience uint
}

type SortWorkers struct {
}

func (a *AddWorkerInfo) addWorkerInfo([]string) {

}

func main() {
	workers := []string{Name: "Михаил", salary: 12000, position: "директор", experience: 5, Name: "Андрей", salary: 11800, position: "мастер", Name: "Игорь", salary: 11000, position: "зам. директора"}

}
