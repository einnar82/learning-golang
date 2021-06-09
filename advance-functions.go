package main

import "fmt"

// you can pass functions as parameters as well
func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

// you can use function closures
// just like PHP closures
func returnFunc(x string) func() {
	return func() {
		fmt.Println("Returns a function", x)
	}
}

// func test() {
// 	fmt.Println("Hello")
// }

// you can assign function
// as a variable  by using
// :=
func main() {
	// x := test
	// x()

	// you can do anonymous functions in go
	// simply declare it as a variable then
	// add func() {} as a value.
	// test := func() {
	// 	fmt.Println("test")
	// }
	// test()

	// you can pass the parameter directly
	// in anonymous functions
	// test := func(x int) int {
	// 	return x
	// }

	// test2(test)
	// fmt.Println(test(1))

	// you can assign a function directly
	// like IIFE in ES6
	// func() {
	// 	fmt.Println("yay")
	// }()

	returnFunc("ge")()
}
