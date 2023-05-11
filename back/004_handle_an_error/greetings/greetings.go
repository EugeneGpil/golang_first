package greetings

import "fmt"
import "errors"

// Hello return the greetings for the named person
func Hello(name string) (string, error) {
	if (name == "") {
		return "", errors.New("empty name")
	}

	// Return a greetings that embeds the name in a message

	// short way
	// message := fmt.Sprintf("Hi %v. Welcome!", name)

	// long way
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
