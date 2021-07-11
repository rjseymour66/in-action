package main

import (
	"log"
	"sync"
	"time"

	"github.com/rjseymour66/in-action/concurrency/patterns/work"
)

// names provides a set of names to display
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason"
}

// namePrinter implements the Worker interface
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	// create a work pool with 2 goroutines.
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names
		for _, name := range names {
			// create a namePrinter and provide the specific name
			np := namePrinter {
				name: name
			}

			go func() {
				// submit the task to be worked on. When RunTask
				// returns we know it is being handled
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wq.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed
	p.Shutdown()
}
