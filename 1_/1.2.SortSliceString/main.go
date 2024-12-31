package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	//sort.Sort(sort.StringSlice(names))
	/*sort.Slice(names, func(i1, i2 int) bool {
		return len(names[i1]) > len(names[i2])
	})*/
	sort.Strings(names)

	fmt.Println(names)
}

func main() {
	names := []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"}
	SortNames(names)
}
