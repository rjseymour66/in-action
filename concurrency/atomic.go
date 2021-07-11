package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter incremented by all goroutines
	counter int64

	// wg waits for the program to finish
	wg sync.WaitGroup
)

func main() {
	// add the goroutines to the WaitGroup
	wg.Add(2)

	// create two goroutines
	go incCounter(1)
	go incCounter(2)

	// make wait for goroutines to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// safely add one to counter
		atomic.AddInt64(&counter, 1)

		// yield thread and get placed back in queue
		runtime.Gosched()
	}
}
