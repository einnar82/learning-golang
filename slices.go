package main

import "fmt"

func main() {
	// array has a definite number of index
	// while slices are doesn't have exact number of index
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3]
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// another way
	// a := []int{5, 6, 7, 8, 9}
	// fmt.Println(cap(a[:3]))
	// append to the end of slice
	// b := append(a, 10)
	// fmt.Println(b)

	// create slice using make()
	newSlice := make([]int, 5)
	fmt.Printf("%T", newSlice)
}
