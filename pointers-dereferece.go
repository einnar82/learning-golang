package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	// & ->  pointer
	// * ->  dereference
	// x := 7
	// y := &x
	// fmt.Println(x, y)
	// *y = 8
	// fmt.Println(x, y)
	toChange := "hello"
	// fmt.Println(toChange)
	// changeValue(&toChange)
	// fmt.Println(toChange)
	var pointer *string = &toChange
	fmt.Println(*pointer)

}
