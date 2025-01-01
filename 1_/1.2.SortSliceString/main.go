package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	// Пример 1
	//sort.Sort(sort.StringSlice(names))

	//Пример 2
	/*sort.Slice(names, func(i1, i2 int) bool {
		return len(names[i1]) > len(names[i2])
	})*/

	//Пример 3
	/*sort.Strings(names)
	j := len(names)
	for i := range names {
		fmt.Print(names[i])
		j--
		if j > 0 {
			fmt.Println(",")
		}
	}*/

	// Пример 4
	sort.Slice(names, func(i, j int) bool {
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
	}
}

func main() {
	names := []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"}
	SortNames(names)
}
