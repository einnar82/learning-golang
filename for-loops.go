package main

import "fmt"

func main() {
	// for loop
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	x := 0
	for x < 5 {
		fmt.Println("hey")
		x++
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 && i != 0 {
			fmt.Println("Divisible by 2")
			continue
		}
	}
}
