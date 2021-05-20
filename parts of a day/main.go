package main

import (
	"time"
)

func main() {

	switch h := time.Now().Hour(); {
	case h <= 7:
		println("good night")
	case h <= 14:
		println("good day")
	default:
		println("good evening")
	}

}
