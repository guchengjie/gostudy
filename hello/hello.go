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

	// 声明一个names的数组
	names := []string{"zhangsan", "lishi"}

	messages, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
