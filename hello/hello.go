package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	/** set properties of the predefined logger **/
	log.SetPrefix("greetings: \n")
	log.SetFlags(0)

	//Request a greeting
	sliceOfNames := []string{"Mothibi", "Thabisile"}
	msg, err := greetings.Hello(sliceOfNames)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
