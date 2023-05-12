package greetings

import "errors"
import "fmt"
import "math/rand"
import "time"

// Hello return the greetings for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greetings that embeds the name in a message

	// short way
	// message := fmt.Sprintf("Hi %v. Welcome!", name)

	// long way
	var message string
	message = fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	formatIndex:=rand.Intn(len(formats))
	return formats[formatIndex]
}
