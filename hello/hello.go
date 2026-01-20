package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	msg := greetings.Hellos([]string{"Mothibi", "Thabisile"})
	fmt.Println(msg)
}
