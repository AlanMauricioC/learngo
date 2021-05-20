package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	word := os.Args[1]
	uperCassedWord := strings.ToUpper(word)
	fmt.Println(uperCassedWord + strings.Repeat("!", len(word)))
}
