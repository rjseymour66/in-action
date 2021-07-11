package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// r is the response, and r.Body is an io.Reader
	r, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	// Create a file to persist the response
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// MultiWriter to write to stdout and a file simultaneously
	dest := io.MultiWriter(os.Stdout, file)

	// Read response and write to both locations
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
