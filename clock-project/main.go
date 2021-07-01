package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	for {

		t := time.Now()

		h, m, s := t.Hour(), t.Minute(), t.Second()

		h1 := h / 10
		h2 := h % 10 //int(math.Mod(float64(h), 10))
		m1 := m / 10
		m2 := m % 10 //int(math.Mod(float64(m), 10))
		s1 := s / 10
		s2 := s % 10 //int(math.Mod(float64(s), 10))

		var sep int
		if s2%2 == 0 {
			sep = 1
		}

		clock := [...]clNumb{
			digits[h1], digits[h2],
			separator[sep],
			digits[m1], digits[m2],
			separator[sep],
			digits[s1], digits[s2],
		}
		for line := range digits[0] {

			for digit := range clock {
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println()
		}

		time.Sleep(time.Millisecond)
		screen.Clear()
		screen.MoveTopLeft()
	}
}
