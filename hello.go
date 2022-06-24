package main

import (
	"github.com/rogery1999/go-greetings-tutorial"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.HelloNames("Rogery", "Valeria")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)
}
