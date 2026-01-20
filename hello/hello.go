package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg, err := greetings.Hello("")
	if err == nil {
		fmt.Println(msg)
	}
	log.Fatal(err)
}
