package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	a := make([]int, len(nums1))
	b := make([]int, len(nums2))
	fmt.Println(len(a), cap(a), len(b), cap(b))
	c := append(a, b...)
	fmt.Println(len(a), len(b))

	return c
}

/*func main() {
	var nums1 []int = []int{1, 2, 3}
	var nums2 []int = []int{}

	fmt.Println(Join(nums1, nums2))
}*/
