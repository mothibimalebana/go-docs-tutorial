// Echo1 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	seperator := " "
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, seperator+os.Args[i])
	}
}
