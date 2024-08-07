package main

import (
	"testing"
)

// Tests creating a new dice object
func TestNewDice(t *testing.T) {
	dice := newDice()

	if len(dice) != 6 {
		t.Errorf("Expected dice length to be 6 but got %d instead", len(dice))
	}
}

//THOUGHT: Maybe test rolling a dice, woudl require mocking a dice roll however?
