package main

import (
	"fmt"
	"sort"
)

func SortNames(names []string) {
	// Пример 1
	/*sort.Sort(sort.StringSlice(names))
	b, _ := json.Marshal(names)
	//fmt.Printf("%v", string(b))
	fmt.Printf("%v", strings.Trim(fmt.Sprintf("%v", string(b)), "[]"))*/

	//Пример 2
	/*sort.Slice(names, func(i, j int) bool {
		return len(names[i]) < len(names[j])
	})
	fmt.Print(names)*/

	//Пример 3
	/*sort.Strings(names)
	var user []string = names
	user = names[:len(names)-1]
	fmt.Print(user)*/
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
	name := ""
	for _, v := range names {

			j--
			if j > 0 {
				fmt.Print(", ")
			} else {
				fmt.Println("")
		}
		name += v
		if v != "" {
			name += ", "
		}
	}
	fmt.Print(name)*/

	// Пример 5
	name := []string{}
	j := 0
	sort.Sort(sort.StringSlice(names))
	// prepend single quote, perform joins, append single quote
	//output := "" + strings.Join(names, `, `) + ``
	for _, v := range names {
		if j < len(names)-1 {
			name = append(name, v+", ")
		} else {
			name = append(name, v)
		}
		j++
	}

	fmt.Println(name)
}

func main() {
	names := []string{"Влада", "Ярослава", "Арина", "Варвара", "Валентина", "Аксинья", "Жанна"}
	SortNames(names)

	/*
		// исходный массив
		initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
		users1 := initialUsers[2:6] // с 3-го по 6-й
		users2 := initialUsers[:4]  // с 1-го по 4-й
		users3 := initialUsers[3:]  // с 4-го до конца

		fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
		fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
		fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]
	*/
}
