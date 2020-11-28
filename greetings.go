package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//Hello returs a greeting for the named personconst Component = require('@linkapi.solutions/nodejs-sdk/component');
func Hello(name string) (string, error) {
	// If no name was given, return an error with message
	if name == "" {
		return "", errors.New("emptys name")
	}

	// Create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos return a map that associates each of the people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {

	// a Map to associete name with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returnde
// message is selected at random.
func randomFormat() string {
	// a Slice of messages formats.
	formats := []string{
		"Hi, %v, Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
