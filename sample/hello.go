package main

import (
	"fmt"
	"log"

	greetings "github.com/VictorAlmeidaFonseca/greetingsEveryOne"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line  number
	log.SetPrefix("greetings: ")
	log.SetFlags(1 | 2)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "Juca"}

	// Get a greeting message and print it
	message, err := greetings.Hellos(names)
	//If an error was returned, print it to the console
	//exit program.
	if err != nil {
		log.Fatal(err)
	}

	// If no errror was returned, print the returned message
	// to the console1
	fmt.Println(message)
}
