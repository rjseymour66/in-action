package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
Goroutines need the following:
	1. Logical processor
	2. WaitGroup so that the main knows to wait for it to finish
	3. Add the number of goroutines to the WG with
	4. goroutine as anonymous function
		- call Done() at the start of the function
	5. Call Wait so main doesn't finish before the goroutines
*/
func main() {
	// Create 1 logical processor
	runtime.GOMAXPROCS(1)

	// Create WaitGroups so that main knows to wait to finish
	// wg is a semaphore -- global var that controls common process
	//   resources
	var wg sync.WaitGroup
	// Add the number of goroutines to the WaitGroup
	wg.Add(2)

	fmt.Println("goroutines start")

	// Create goroutine with anonymous function
	go func() {
		// defer to make sure to tell main its done
		defer wg.Done()

		// Count to 26
		for count := 1; count < 27; count++ {
			fmt.Printf("%d ", count)
		}
	}()

	go func() {
		defer wg.Done()

		for c := 'A'; c < 'A'+26; c++ {
			fmt.Printf("%c ", c)
		}
	}()

	// Wait for goroutine to finish
	fmt.Println("Waiting for goroutines")
	wg.Wait()

	fmt.Println("\nDone")
}
