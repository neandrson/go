package main

import "fmt"

func Join(nums1, nums2 []int) []int {
	fmt.Println(cap(nums1), cap(nums2))
	a := make([]int, cap(nums1)+cap(nums2))
	//fmt.Println(cap(a))
	a = append(nums1)
	//copy(a, nums2)
	fmt.Println(a, len(a), cap(a))
	//a = make([]a, len(nums1)+len(nums2))
	//fmt.Println(a, len(a), cap(a))
	//j := 0
	/*for i := 0; i <= cap(nums1); i++ {
		a = append(a, nums1[i])
		fmt.Println(a, cap(a))
	}*/
	for i := 0; i <= cap(nums2); i++ {
		a = append(a, nums2[i])
		fmt.Println(a, cap(a))
	}
	//copy(a, a[:len(nums1)+len(nums2)])
	//a = a[len(nums1)+len(nums2) : len(nums1)+len(nums2)]
	fmt.Println(a)

	return a
}

/*func main() {
	var nums1 []int = []int{1, 2, 3}
	var nums2 []int = []int{}

	fmt.Println(Join(nums1, nums2))
}*/
