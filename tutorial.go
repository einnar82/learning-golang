package main

import "fmt"

func main()  {
	// concatinate string in go
	var name string = "rannie"
	result := fmt.Sprintf("%s %s", "Hello", name)
	fmt.Println(result)
	// another one
	var lastname string = "again"
	fmt.Println("Hello " + lastname)
}