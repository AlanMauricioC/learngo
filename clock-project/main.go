package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

type clNumb = [5]string

func main() {
	digits := [10]clNumb{
		{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"  █",
			"  █",
		},
	}

	separator := [2]clNumb{
		{
			"   ",
			" ░ ",
			"   ",
			" ░ ",
			"   ",
		},
		{
			"   ",
			"   ",
			"   ",
			"   ",
			"   ",
		},
	}
	for {

		for i := 0; i < 5; i++ {
			t := time.Now()

			h := t.Hour()
			m := t.Minute()
			s := t.Second()

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
			fmt.Printf("%s %s %s %s %s %s %s %s \n", digits[h1][i], digits[h2][i], separator[sep][i], digits[m1][i], digits[m2][i], separator[sep][i], digits[s1][i], digits[s2][i])
		}
		time.Sleep(time.Millisecond)
		screen.Clear()
		screen.MoveTopLeft()
	}
}
