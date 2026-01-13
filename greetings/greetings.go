package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi %s, hope you brought pizza", name)
	return message
}
