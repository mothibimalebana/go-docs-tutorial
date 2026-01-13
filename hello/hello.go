package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	msg := greetings.Hello("Mothibi")
	fmt.Println(msg)
}
