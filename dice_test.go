package main

import (
	"testing"
)

func TestNewDice(t *testing.T) {
	dice := newDice()

	if len(dice) != 6 {
		t.Errorf("Expected dice length to be 6 but got %d instead", len(dice))
	}
}
