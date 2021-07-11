/* A runner is useful when developing a program that is scheduled to run as a 
   background process.

   The runner terminates when the following events occur:
	- The program completes its work within the time alloted and terminates normally
	- The program cannot complete its work in time and kills itself
	- An os interrupt event is received and the program tries to shut down cleanly
*/
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can shut down on 
// an operating system interrupt
type Runner struct {
	// reports a signal from the os
	interrupt chan os.Signal

	// reports that processing is done
	complete chan error

	// reports that time has run out
	timeout <-chan time.Time

	// holds a set of functions that are executed
	// as a series in index order
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the timeout
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt is returned when an event from the os is received
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready-to-use Runner.
// Each channel field is initialized. tasks is a nil slice.
func New(d time.Duration) *Runner {
	return &Runner {
		interrupt:	make(chan os.Signal, 1),
		complete:	make(chan error),
		timeout:	time.After(d),
	}
}

// Add attaches tasks to the Runner. A task is a function that
// takes an int ID.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	// We want to receive all interrupt based signals
	signal.Notify(r.interrupt, os.Interrupt)

	// Run each task on its own goroutine
	go func() {
		r.complete <- r.run()
	}()

	select {
	// Signaled when processing is done
	case err := <-r.complete:
		return err

	// Signaled when time runs out
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each registered task
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// Check for an interrupt signal from the OS.
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// Execute the registered task
		task(id)
	}
	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
	select {

	// Signaled when an interrupt event is sent.
	case <-r.interrupt:
		// Stop receiving any further signals
		signal.Stop(r.interrupt)
		return true

	// Continue running as normal
	default:
		return false
	}
}

