package main

import (
	"log"
)

// set log in the init so that it can be used as soon as the program starts
func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger
	log.Println("message")

	// Fatalln is Println() followed by os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println followed by a call to panic()
	log.Panicln("panic message")
}
