package main

import (
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {

}

func main() {
	var b, c int
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Scan("Введите длину и объём массива ", &b, &c)
	d, err := fmt.Println(UnderLimit(a, b, c))
	if err != nil {
		fmt.Println("Превышен объём")
	}
	fmt.Println(d)
}
