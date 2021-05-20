package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage [name] [positive|negative]")
		return
	}
	name, mType := os.Args[1], os.Args[2]

	if name == "" {
		fmt.Println("Please provide [your name]")
		return
	}

	if mType == "" {
		fmt.Println("Please provide [positive|negative]")
		return
	}
	moods := [...][3]string{
		{
			"really cool",
			"good",
			"Happy",
		}, {
			"bad",
			"sad :c",
			"hungry",
		},
	}

	var t int
	if mType != "positive" {
		t = 1
	}

	s := time.Now().UnixNano()
	rand.Seed(s)

	const lmood = len(moods[t])
	mood := moods[t][rand.Intn(lmood)]

	fmt.Printf("%v feels %v ", name, mood)
}
