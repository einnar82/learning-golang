package main

import (
	"fmt"
	"sync"
)

// mutex and waitgroup are partners for goroutines
// waitgroup is for wait the sequence
// mutex is to lock the sequence or syncronize the pattern
var waitGroup = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

func main() {
	// create OS thread
	// runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		waitGroup.Add(2)
		mutex.RLock()
		go sayHello()
		mutex.Lock()
		go increment()
	}
	waitGroup.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	mutex.RUnlock()
	waitGroup.Done()
}

func increment() {
	counter++
	mutex.Unlock()
	waitGroup.Done()
}
