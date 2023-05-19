package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting to the user.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// Hellos returns greetings for multiple names.
func Hellos(names []string) (map[string]string, error) {
	// The map of the names to their greetings
	messages := make(map[string]string, len(names))
	for _, name := range names {
		personalizedGreeting, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = personalizedGreeting
	}
	return messages, nil
}

// init sets initial values of the variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat selects and returns a welcome message at random
func randomFormat() string {
	// 'formats' defines the list of messages to select from
	formats := []string{
		"Hello, %v. Hope you brought pizza!",
		"Welcome to the club, %v. It's great to have you here!",
		"%v just showed up; the room just got brighter.",
	}
	return formats[rand.Intn(len(formats))]
}
