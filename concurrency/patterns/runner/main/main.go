package main

import (
	"log"
	"os"
	"time"

	"github.com/rjseymour66/in-action/concurrency/patterns/runner"
)


// timeout is the number of seconds the program has to finish
const timeout = 3 * time.Second

// main is the entry point for the program
func main() {
	log.Println("Starting work.")

	// Create a new timer value for this run
	r := runner.New(timeout)

	// Add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// Run the tasks and handle the result
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
