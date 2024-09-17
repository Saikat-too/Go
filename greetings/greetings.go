package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// if no name was given , return  an error with message .
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf("Hi , %v. Welcome! ", name)
	return message, nil
}
