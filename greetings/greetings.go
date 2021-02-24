package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("The name can't be blank!")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// take an array
// returns map[string]string and error
func Greetings(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {

		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// return a random string
func randomFormat() string {
	formats := []string{
		"Hi %v. How are you?",
		"Merhaba %v. Nasılsın?",
		"Hi. I'm %v. Nice to meet you",
	}

	return formats[rand.Intn(len(formats))]
}
