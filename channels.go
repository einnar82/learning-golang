package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	wg.Add(2)
	// receiving goroutine
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println("receive")
		fmt.Println(i)
		wg.Done()
	}(ch)
	// sending goroutine
	go func(ch chan<- int) {	
		ch <- 42
		fmt.Println("send")
		wg.Done()
	}(ch)

	wg.Wait()
}
