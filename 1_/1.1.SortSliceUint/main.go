package main

import (
	"fmt"
	"slices"
)

func SortNums(nums []uint) {
	//slices.Sort(nums)
	slices.SortFunc(nums, func(a, b uint) int {
		// Здесь мы сравниваем сами
		switch {
		case a < b:
			return -1
		case a > b:
			return 1
		default:
			return 0
		}
	})
	fmt.Println(nums)
}

func main() {
	smallInts := []uint{490, 741, 88, 1, 10, 7, 234, 2234, 64, 12, 778, 21234, 4345, 45673, 23, 5, 78, 2, 1, 5}
	SortNums(smallInts)
}
