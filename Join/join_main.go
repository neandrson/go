package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	a := make([]int, len(nums1)+len(nums2))
	a = append(nums1, nums2...)
	copy(a, a[:len(nums1)+len(nums2)])
	//a = a[len(nums1)+len(nums2) : len(nums1)+len(nums2)]
	fmt.Println(len(a), cap(a), a)

	return a
}

/*func main() {
	var nums1 []int = []int{1, 2, 3}
	var nums2 []int = []int{}

	fmt.Println(Join(nums1, nums2))
}*/
