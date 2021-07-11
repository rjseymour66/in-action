package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// create the channel
	baton := make(chan int)

	// Add to WG
	wg.Add(2)

	// First runner starts
	go Runner(baton)

	// Start race
	baton <- 1

	// Wait for goroutine
	wg.Wait()

}

// Simulates a person running in the relay race
func Runner(baton chan int) {

	var newRunner int

	// Wait to receive the baton from the channel
	runner := <-baton

	// Start running
	fmt.Printf("Runner %d Running with Baton\n", runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the Line\n", newRunner)
		go Runner(baton)
	}

	// sleep represents running around the track
	time.Sleep(100 * time.Millisecond)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton
	fmt.Printf("Runner %d Exchange with Runner %d\n", runner, newRunner)

	baton <- newRunner
}
