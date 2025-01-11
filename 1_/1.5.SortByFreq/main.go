package main

import (
	"fmt"
	"sort"
)

type item struct {
	char      byte
	frequency int
}

func SortByFreq(s string) string {
	stringMap := make(map[byte]int)
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		stringMap[s[i]]++
	}
	itemArray := make([]item, 0)
	for key, value := range stringMap {
		i := item{char: key, frequency: value}
		itemArray = append(itemArray, i)
	}
	sort.Slice(itemArray, func(i, j int) bool {
		return itemArray[i].char < itemArray[j].char
	})
	sort.Slice(itemArray, func(i, j int) bool {
		return itemArray[i].frequency < itemArray[j].frequency
	})
	output := ""
	for i := 0; i < len(itemArray); i++ {
		for j := 0; j < itemArray[i].frequency; j++ {
			output = output + string(itemArray[i].char)
		}
	}
	return output
}

func main() {
	output := SortByFreq("zggoooaaarrygyzv")
	fmt.Println(output)
}
