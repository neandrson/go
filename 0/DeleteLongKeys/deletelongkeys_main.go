package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	myMap := make(map[string]int)
	num := 6
	for key, _ := range m {
		if len(key) < num {
			fmt.Println(m)
			delete(key, m)
			myMap = m
		}
	}
	return myMap
}

func main() {
	m := map[string]int{"aaaaaaa": 1, "bb": 2, "cccccc": 3}
	fmt.Println(DeleteLongKeys(m))
}
