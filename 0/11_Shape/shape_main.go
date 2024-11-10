package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	pi float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	if r.Height > 0 && r.Width > 0 {
		return r.Height * r.Width
	}
	return -1
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	figure_1 := Circle{radius: 1.0}
	fmt.Println(figure_1.Area())
	figure_2 := Rectangle{width: 57.2, height: 10.2}
	fmt.Println(figure_2.Area())
}
