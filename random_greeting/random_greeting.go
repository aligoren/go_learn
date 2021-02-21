package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("empty name!")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func init() {

	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {

	formats := []string{
		"Hi, %v welcome. How are you today?",
		"I love you %v",
		"Do you love me %v?",
	}

	return formats[rand.Intn(len(formats))]
}

func main() {

	log.SetPrefix("Random greetings: ")
	log.SetFlags(0)

	message, err := Hello("Ali")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
