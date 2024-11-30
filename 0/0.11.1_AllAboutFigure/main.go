package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	pi     float64
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Circle) Area() float64 {
	r.pi = 3.141592653589793
	return r.pi * (r.radius * r.radius)
}

func (r Rectangle) Area() float64 {
	if r.height > 0 && r.width > 0 {
		return r.height * r.width
	}
	return -1
}

/*func GetArea(s Shape) float64 {
	return s.Area()
}*/

func main() {
	figure_1 := Circle{radius: 1.0}
	fmt.Println(figure_1.Area())
	figure_2 := Rectangle{width: 57.2, height: 10.2}
	fmt.Println(figure_2.Area())
}
