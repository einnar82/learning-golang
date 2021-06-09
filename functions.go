package main

import "fmt"

// func test(x int) {
// 	fmt.Println("test", x)
// }

// to add data type
// varible dataType
// for multiple variables with same data type
// you can use variableA, variableB type
func add(x, y int) int {
	return x + y
}

// you can return multple values
// and types
// then destructure the function
// func subAndTimes(x, y int) (int, int) {
// 	return x - y, x * y
// }

// labeling the return types
// you can label your return types for
// code readability.
func subAndTimes(x, y int) (subtract, multiply int) {
	// A defer statement postpones the execution
	// of a function until the surrounding function returns,
	// either normally or through a panic.
	defer fmt.Println("Done computing..")
	subtract = x - y
	multiply = x * y
	fmt.Println("Before computing..")
	return
}

func main() {
	// test(1)
	// ans := add(1, 2)
	difference, product := subAndTimes(10, 2)
	// fmt.Println(ans)
	fmt.Println(difference)
	fmt.Println(product)
}
