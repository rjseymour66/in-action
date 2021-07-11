package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // number of goroutines to use
	taskLoad	 = 10 // amount of work to process
)

var wg sync.WaitGroup

// initialize the package by Go runtime prior to any other code execution
func init() {
	// seed random number generator
	rand.Seed(time.Now().Unix())
}

func main() {
	// create buffered chan for task load
	tasks := make(chan string, taskLoad)

	// launch goroutines to do work
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// close the channel, making all goroutines quit
	// when work is done
	close(tasks)

	wg.Wait()
}

// worker is launched as a goroutine to process work from the buffered channel
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// While there are still tasks in the buffer, 
		// wait for tasks to be assigned
		task, ok := <-tasks
		if !ok {
			// The channel is empty and closed
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// Display that we are starting the work
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// Randomly wait to simulate work time
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display work is finished
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
