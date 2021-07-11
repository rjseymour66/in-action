package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter incremented by all goroutines
	counter int

	// wg waits for the program to finish
	wg sync.WaitGroup

	// mutex is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	// add the goroutines to the WaitGroup
	wg.Add(2)

	// create two goroutines
	go incCounter(1)
	go incCounter(2)

	// make wait for goroutines to finish
	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()

		// capture value of counter
		value := counter

		// yield the thread and be placed back in the queue
		runtime.Gosched()

		// increment local value of counter
		value++

		// store value back into counter
		counter = value

		mutex.Unlock()
	}
}
