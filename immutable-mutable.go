package main

import "fmt"

func main() {
	// slices, maps are mutable
	var x int = 5
	// just like pointers
	// can store and manipulate data
	y := x
	y = 7
	fmt.Println(x, y)
}
