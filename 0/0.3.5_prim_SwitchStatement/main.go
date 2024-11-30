package main

import "fmt"

func main() {
	/*weather := "rain"
	switch weather {
	case "rain":
		fmt.Println("Погода дождливая")
	case "sunny":
		fmt.Println("Погода солнечная")
	default:
		fmt.Println("Погода не определена")
	}*/

	/*switch variable {
	case value1:
		// выполняемые действия, если variable равна value1
	case value2:
		// выполняемые действия, если variable равна value2
	...
	default:
		// выполняемые действия, если variable не равна ни одному из значений
	}*/

	var i int
	switch {
	case i > 0:
		fmt.Println("i больше нуля")
	case i < 0:
		fmt.Println("i меньше нуля")
	default:
		fmt.Println("i равно нулю")
	}
}
