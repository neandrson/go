package main

import "fmt"

func DeleteLongKeys(m map[string]int) map[string]int {
	myMap := make(map[string]int)
	fmt.Println(m)
	for str, _ := range m {
		if len(str) >= 6 {
			myMap[str]++
		}
	}
	return myMap
}

func main() {
	m := map[string]int{"aaaaaaa": 1, "bb": 2, "cccccc": 3}
	fmt.Println(DeleteLongKeys(m))
}
