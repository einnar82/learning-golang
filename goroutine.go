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

// use sync.WaitGroup to wait groups of
// goroutine to complete

// use sync.Mutex and sync.RWMutex to protect data access

/**
*	Parallelism, by default Go will use CPU threads equal to available cores.
*	You can change it via runtime.GOMAXPROCS
*   More threads can increase performance, but too many can slow it down.
* 	Let consumer control concurrency.
*
* 	when creating a goroutine, know how it will end.
* 	Avoid subtle memory leaks.
* 	Check for race conditions at compile time.
**/
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
