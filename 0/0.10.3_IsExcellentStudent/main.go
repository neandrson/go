package main

import "fmt"

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	b := false
	if float64(s.solvedProblems)*s.scoreForOneTask >= s.passingScore {
		b = true
	}
	return b
}

func main() {
	student := Student{name: "Arseniy", solvedProblems: 5, scoreForOneTask: 1.0, passingScore: 100.0}
	fmt.Print(student.IsExcellentStudent())
}
