// Dup2: print the names of all files in which each dup lic ated line occ urs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	fileNames := os.Args[1:]

	countLines(fileNames, counter)

}

func countLines(fileNames []string, counter map[string]int) {

	if len(fileNames) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			counter[input.Text()]++
		}
		for duplicate_word, instances := range counter {
			fmt.Printf("%d\t%s\n", instances, duplicate_word)
		}
	} else {
		for index, file := range fileNames {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				continue
			} else {
				input := bufio.NewScanner(f)
				for input.Scan() {
					counter[input.Text()]++
				}

				fmt.Print(os.Args[index+1], "\n")
				for duplicate_word, instances := range counter {
					fmt.Printf("%d\t%s\n", instances, duplicate_word)
				}
			}
		}
	}

}
