// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v", err)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read: %v", err)
		}
		fmt.Printf("<------------Body Section-------->\n")
		fmt.Printf("%v\n", b)
		fmt.Printf("<------------Header Section-------->\n")
		fmt.Printf("%v\n", resp.Header)
	}
}
