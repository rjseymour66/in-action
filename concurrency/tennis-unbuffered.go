package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// create an unbuffered channel
	court := make(chan int)

	// tell main to wait for 2 goroutines
	wg.Add(2)

	// Launch the players
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// Use v, ok to test if there is a ball in the buf
		ball, ok := <-court
		if !ok {
			// if channel was closed we won
			fmt.Printf("Player %s Won!\n", name)
			return
		}

		// Pick random num and see if we miss the ball
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost
			close(court)
			return
		}

		// Display and then increment the hit count by one
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit ball back
		court <- ball
	}
}
