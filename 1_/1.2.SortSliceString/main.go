package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	// Пример 1
	//sort.Sort(sort.StringSlice(names))
	//user := names[:len(names)-1]
	//fmt.Print(names)

	//Пример 2
	/*sort.Slice(names, func(i, j int) bool {
		return len(names[i]) < len(names[j])
	})
	fmt.Print(names)*/

	//Пример 3
	sort.Strings(names)
	var user []string = names
	user = names[:len(names)-1]
	fmt.Print(user)
	/*j := len(names)
	for i := range names {
		fmt.Print(names[i])
		j--
		if j > 0 {
			fmt.Println(",")
		}
	}*/

	// Пример 4
	/*sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	j := len(names)
	for _, v := range names {
		fmt.Print(v)
		j--
		if j > 0 {
			fmt.Print(", ")
		} else {
			fmt.Println("")
		}
	}*/
}

func main() {
	names := []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"}
	SortNames(names)
}
