package main

import "fmt"

type Shape interface {
	Area() float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

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
	//area := rect.Area()
	area := CalculateArea(rect)
	scaledW, scaledH := rect.Scale(2.0)
	fmt.Println(scaledW, scaledH)
	fmt.Println(area)
}
