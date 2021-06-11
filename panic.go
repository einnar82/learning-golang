package main

import "fmt"

func main() {

	// order of execution
	// defer, panic
	defer fmt.Println("After")
	fmt.Println("Before")
	panic("Error")
}
