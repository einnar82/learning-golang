package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	fmt.Println("done panicking")
	panic("\nsomething happened here")

}

func main() {

	// order of execution
	// defer, panic
	// defer fmt.Println("After")
	// fmt.Println("Before")
	// panic("Error")
	fmt.Println("Start")
	panicker()
	fmt.Println("end")
}
