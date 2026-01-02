package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{
		"Bello",
		"Olamilekan",
		"Olayinka",
	}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
