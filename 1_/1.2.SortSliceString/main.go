package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	sort.Sort(sort.StringSlice(names))
	fmt.Println(names)
}

func main() {
	names := []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"}
	SortNames(names)
}
