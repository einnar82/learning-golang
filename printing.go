package main 

import "fmt"

func main() {
	// print in console
	// %T = data type
	// %v = value of variable
	fmt.Printf("hello %T %v", 10, 10)
	
	var out string = fmt.Sprintf("number: \t %07d is cool", 45)
	fmt.Println(out)
}