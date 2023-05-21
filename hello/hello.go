package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)
	

func main() {
	log.SetPrefix("greetins: ")
	log.SetFlags(0)

	names := []string{"Gustavo", "Clarisse", "Angela"}

	messages, err := greetings.Hellos(names) 
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}