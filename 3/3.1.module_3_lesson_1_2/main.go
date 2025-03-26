package main

type MyContains interface {
	int | float64
}

func Filter[T MyContains](nums []T) T {
	var total T
	s, ok := a.(int)
	if ok {
		fmt.Printf("'%v' is a string\n", s)
	} else {
		fmt.Printf("'%v' is not a string\n", a)
	}
}

func main() {

}
