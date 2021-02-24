package main

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name!")
	}

	message := fmt.Sprintf("Hi, %v. How are you?", name)

	return message, nil
}

func main() {
	log.SetPrefix("Error Handling: ")
	log.SetFlags(0)

	message, err := Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
