package main

import "fmt"

func main() {
	var firstName string = "John"
	lastName := "Doe"
	fullName := fmt.Sprintf("%v %v", firstName ,lastName)
	fmt.Println(fullName)
}