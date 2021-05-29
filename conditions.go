package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6.5
	name1 := "a"
	name2 := "b"
	val := float64(x) != y
	name := name1 < name2
	fmt.Printf("%t", val)
	fmt.Printf("%t", name)
	// note:
	// when comparing variables, 
	// it must be same datatypes in able to perform comparisons.
}