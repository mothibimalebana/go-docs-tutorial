package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//list of possible messages to send
	var err error
	var msg string

	messageSlice := []string{
		"Hi, %s we hope I hope you brought pizza",
		"%s just crash landed",
		"Welcome %s",
		"%s is here",
	}

	if name == "" {
		err = errors.New("entered an empty string")
		return "", err
	}
	msgIndex := rand.Intn(len(messageSlice))
	msg = fmt.Sprintf(messageSlice[msgIndex], name)
	return msg, nil
}

func Hellos(names []string) (map[string]string, error) {

	greetingsMap := make(map[string]string)
	var err error

	for _, name := range names {

		greetingsMap[name], err = Hello(name)
	}

	return greetingsMap, err
}
