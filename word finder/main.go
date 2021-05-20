package main

import (
	"fmt"
	"os"
	"strings"
)

const text = "this is a text about how how the text works"

func main() {
	query := os.Args[0:]
	words := strings.Fields(text)

	for _, q := range query {
		for _, w := range words {
			if q == w {
				fmt.Printf("found %v", w)
				break
			}
		}
	}
}
