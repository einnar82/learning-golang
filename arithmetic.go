package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int = 9
	var num2 int = 4

	answer := num1 % num2
	pi := math.Pi
	fmt.Printf("%d", answer)
	fmt.Println(pi)
	// BEDMAS in Go lang
	// Brackets, exponents, division, multiplication, addition, subtraction 
}