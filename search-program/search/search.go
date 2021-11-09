package search

import (
	"log"
	"sync"
)

// a map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results
	restuls := make(chan *Result)

	// Setup a waitgroup so we can process all of the feeds
	var waitGroup sync.WaitGroup

	// set the number of goroutines we need to wait for while
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
	}
}
