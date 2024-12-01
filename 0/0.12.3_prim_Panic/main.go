package main

import "fmt"

/*func main() {
	fmt.Println("Начало")
	panic("Что-то пошло не так")
	fmt.Println("Конец")
}*/

/*func exampleFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	// какой-то код, в котором может возникнуть паника
}*/

func main() {
	fmt.Println("Start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("Something went wrong")
	fmt.Println("End")
}
