package main

import "fmt"

func main() {
	// normal switch
	ans := 10
	switch ans {
	case 1:
		fmt.Println("answer is ", ans)
	case 2, 9:
		fmt.Println("answer is ", ans)
	default:
		fmt.Println("answer is 10")
	}

	// switch without variable declaration
	switch {
	case ans == 10:
		fmt.Println("Yeah it's 10")
	default:
		fmt.Printf("Oh no!")
	}
}
