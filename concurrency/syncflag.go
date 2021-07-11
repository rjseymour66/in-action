package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// alerts running goroutines to shutdown
	shutdown int64

	// create WaitGroup
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(5 * time.Second)

	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// wait for the goroutines to finish
	wg.Wait()
}

// doWork checks the shutdown flag to terminate early
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown?
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
