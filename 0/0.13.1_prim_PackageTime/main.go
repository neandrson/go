package main

import (
	"fmt"
	"time"
)

func main() {
	//now := time.Now()
	//fmt.Println("Current time:", now)

	t1 := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println("Difference between t1 and t2:", diff)
}
