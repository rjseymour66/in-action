package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to Finish")

	wg.Wait()
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer % inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
