package main

import (
	"fmt"
	"log"

	"github.com/aligoren/go_learn/greetings"
)

func main() {

	names := []string{"Ali", "Burak", "Can"}

	messages, err := greetings.Greetings(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
