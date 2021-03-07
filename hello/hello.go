package main

import (
	"fmt"
	"log"
	"mrgu/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("fsdgsfd")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
}
