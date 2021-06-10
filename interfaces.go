package main

import (
	"fmt"
	"math"
)

// Interface is just like the interface implementation
// in other programming language like PHP, JS

//
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	r1 := Rectangle{4.3, 4}
	shapes := []Shape{&c1, &r1}

	// shapes[0].area()
	// fmt.Println(shapes[1].area())
	for _, shape := range shapes {
		// fmt.Println(shape.area())
		fmt.Println(getArea(shape))
	}
}
