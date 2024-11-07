package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	myMap := make(map[string]int)
	myMap = m
	for str, _ := range myMap {
		if len(str) < 6 {
			delete(myMap, str)
			continue
		}
	}
	return myMap
}

func main() {
	m := map[string]int{"aaaaaaa": 1, "bb": 2, "cccccc": 3}
	fmt.Println(DeleteLongKeys(m))
}
