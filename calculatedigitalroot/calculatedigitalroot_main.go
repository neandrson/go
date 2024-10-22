package main

import (
	"fmt"
)

var res, cal, sum int

func CalculateDigitalRoot(n int) int {

	if n < 10 {
		return n
	}
	//fmt.Println(n)
	res = (n % 10) + CalculateDigitalRoot(n/10)
	//fmt.Println(res)
	if res >= 10 {
		CalculateDigitalRoot(res)
	}
	/*for ; n > 0;

	return n%10 + CalculateDigitalRoot(n/10)*/
	/*var res, i int
	res = 0
	var a string

	a = strconv.Itoa(n)

	for i = 0; i < len(a); i++ {
		val, err := strconv.Atoi(string(a[i]))
		if err == nil {
			res += val
		}
	}
	if res > 9 {
		return CalculateDigitalRoot(res)
	}*/
	//fmt.Println(n%10, n/10, res)
	return res
	// CalculateDigitalRoot(n%10) + CalculateDigitalRoot(n/10)

}

func main() {
	fmt.Println(CalculateDigitalRoot(12345))
}
