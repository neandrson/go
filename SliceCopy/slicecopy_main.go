package main

import (
	"slices"
)

func SliceCopy(nums []int) []int {
	return slices.Clone(nums)
}

func main() {
	a := []int{}
	SliceCopy(a)
}
