package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Scanln(&age) //Запрос возраста
	if age < 18 {
		fmt.Println("Вы несовершеннолетний")
	} else if age < 65 {
		fmt.Println("Вы взрослый")
	} else {
		fmt.Println("Вы очень взрослый")
	}
}
