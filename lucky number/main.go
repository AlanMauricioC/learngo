package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	s := time.Now().UnixNano()
	rand.Seed(s)

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide a number")
		return
	}
	guess, e := strconv.Atoi(args[0])
	if e != nil {
		fmt.Println("Please provide a valid number")
		return
	}

	if guess < 0 {
		fmt.Println("Please provide a positive number")
		return
	}
	for turn := 0; turn < 10; turn++ {
		if rand.Intn(11) == guess {
			fmt.Println("You win!!")
			return
		}
	}

	fmt.Println("You are a losser! :0")
}
