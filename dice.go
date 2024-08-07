package main

import (
	"fmt"
	"math/rand"
	"time"
)

//THOUGHT: Could have possibly made an interface for dice that made it easier to test?

type dice []string

// Constructor for dice
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

// Rolls dice with random source as current time
func (d dice) roll() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	rn := r.Intn(len(d))

	fmt.Printf("You rolled a %s\n", d[rn])

	return (rn + 1)
}
