package main

import (
	"fmt"
	"sort"
)

/*
	type Chislo struct {
		Left  []int
		Right []int
	}
*/
/*type Num struct {
	left  []int
	right []int
}*/

func SortAndMerge(left, right []int) []int {
	final := []int{}
	sort.Ints(left)
	sort.Ints(right)
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	//sort.Ints(left)
	//sort.Ints(right)
	//final = copy(left, right)

	return final
}

func main() {
	left := []int{490, 741, 88, 1, 10, 7, 234, 2234, 64, -12, 778, 21234, 4345, 45673, -23, 5, 78, 2, 1, 5}
	right := []int{-1, 4, 5, 104, 1, 18733, 0}
	a := SortAndMerge(left, right)

	fmt.Println(a)
}
