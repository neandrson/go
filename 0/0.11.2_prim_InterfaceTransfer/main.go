package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Scale(factor float64) (float64, float64) {
	r.width *= factor
	r.height *= factor
	return r.width, r.height
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	area := rect.Area()
	scaledW, scaledH := rect.Scale(2.0)
	fmt.Println(scaledW, scaledH)
	fmt.Println(area)
}
