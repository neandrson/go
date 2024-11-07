package main

import "fmt"

func CountingSort(contacts []string) map[string]int {
	countMap := make(map[string]int)
	for _, contact := range contacts {
		countMap[contact]++
	}
	return countMap
}

func main() {
	contacts := []string{"Alice", "Bob", "Charlie", "Alice", "Bob", "Alice"}
	result := CountingSort(contacts)
	fmt.Println("Count of contacts:", result)
	fmt.Println("Duplicates:")
	for contact, count := range result {
		if count > 1 {
			fmt.Printf("%s: %d\n", contact, count)
		}
	}
}
