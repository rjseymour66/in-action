package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// create the buffer
	var b bytes.Buffer

	// buffer has Read and Write methods
	b.Write([]byte("Hello "))

	// Fprintf writes to the console
	fmt.Fprintf(&b, "World!\n")

	b.WriteTo(os.Stdout)
}
