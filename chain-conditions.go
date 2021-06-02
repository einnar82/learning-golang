package main

import "fmt"

func main() {
	val := 5 < 7 || 7 > 9
	fmt.Printf("%t", val)
	// ! not
	// || or
	// && and
}
