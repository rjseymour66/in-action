package main

import (
	"log"
	"os"

	// blank identifier allows the compiler to accept the import
	// and call any init functions that are in this package
	_ "github.com/rjseymour66/in-action/search-program/matchers"
	"github.com/rjseymour66/in-action/search-program/search"
)

// init is called prior to main()
func init() {
	// change the device for logging to stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// Perform the search for the specified term
	search.Run("president")
}
