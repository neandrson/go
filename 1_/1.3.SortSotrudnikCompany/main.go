package main

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type AddWorkerInfo struct {
	name, position string
	salary, experience uint
}

type SortWorkers struct {
	
}

func (a *AddWorkerInfo) addWorkerInfo {

}

func main() {

}
