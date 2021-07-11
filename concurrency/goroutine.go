package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all go programs
func main() {
	// Allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(2)

	// wg is used to tell main to wait for the program to finish
	// Add a count of two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine
	go func() {
		// Schedule the call to done to tell main we're done
		// Done decrements the Add(2) WaitGroup
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// Declare an anonymous function and create a goroutine
	go func() {
		// Schedule the call to done to tell main we're done
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// wait for the goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
