package main

import (
	"fmt"
	"math/rand"
	"time"
)

type dice []string

func newDice() dice {
	dice := []string{
		"ONE",
		"TWO",
		"THREE",
		"FOUR",
		"FIVE",
		"SIX",
	}

	return dice
}

func (d dice) roll() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	rn := r.Intn(len(d))

	fmt.Printf("You rolled a %s\n", d[rn])

	return (rn + 1)
}
