package main

import "fmt"

// struct is a custom type in go
type Point struct {
	x        int32
	y        int32
	isOnGrid bool
}

// Embedded struct is combination of 1 or more struct inside of a struct
type Circle struct {
	radius float64
	center *Point
}

func changeX(pt *Point) {
	pt.x = 100
}

func main() {
	// var p1 Point = Point{1, 2, false}
	// p1 := &Point{1, 2, false}
	// var p2 Point = Point{-5, 7, true}
	// p1.x = 7
	// fmt.Println(p1)
	// changeX(p1)
	// c1 := Circle{4.56, p1}
	c1 := Circle{4.56, &Point{1, 1, true}}
	fmt.Println(c1.center.x)
}
