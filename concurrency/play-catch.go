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

	air := make(chan int)

	wg.Add(2)

	// Launch people playing catch
	go player("Father", air)
	go player("Son", air)

	// Throw the first ball
	air <- 1

	// Wait for them to finish playing
	wg.Wait()
}

func player(name string, air chan int) {
	defer wg.Done()

	for {
		// test if someone has a ball
		ball, ok := <- air
		if !ok {
			fmt.Printf("%s doesn't want to play anymore!\n", name)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("%s dropped the ball!\n", name)

			close(air)
			return
		}

		fmt.Printf("%s caught the ball\n", name)
		ball++

		// throw the ball back
		air <- ball
	}
}
