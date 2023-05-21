package main

import (
	"fmt"

	"example.com/greetings"
)
	

func main() {
	message := greetings.Hello("Angela")
	fmt.Println(message)
}