package main

import (
	"fmt"
)

func MinPizzaCost(s, m, l, cs, cm, cl, x int) int {
	scs := float64(s) / float64(cs)
	mcm := float64(m) / float64(cm)
	lcl := float64(l) / float64(cl)

	sum := (scs + mcm + lcl) / 3.0 * float64(x)
	fmt.Println(scs, mcm, lcl, sum)
	return int(sum)
}

func main() {
	s := 314
	m := 706
	l := 1256
	cs := 230
	cm := 510
	cl := 925
	x := 500
	a := MinPizzaCost(s, m, l, cs, cm, cl, x)
	fmt.Println(a)
}
