package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func main() {
	c := NewCircle(5)
	fmt.Println(c.Area())
}
