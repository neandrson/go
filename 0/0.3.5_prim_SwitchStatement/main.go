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

/*package main

import "fmt"

func main() {
    i := 2
    if i == 1 {
        fmt.Println("i равно 1")
    } else if i == 2 {
        fmt.Println("i равно 2")
    } else {
        fmt.Println("i не равно ни 1, ни 2")
    }
}*/

/*package main

import "fmt"

func main() {
    i := 2
    switch i {
    case 1:
        fmt.Println("i равно 1")
    case 2:
        fmt.Println("i равно 2")
    default:
        fmt.Println("i не равно ни 1, ни 2")
    }
}*/
