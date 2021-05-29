package main

import "fmt"

func main() {
	var firstName string = "John"
	lastName := "Doe"
	fullName := fmt.Sprintf("%s %s", firstName ,lastName)
	fmt.Println(fullName)
}