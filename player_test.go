package main

import (
	"testing"
)

// Tests creating a new player object
func TestNewPlayer(t *testing.T) {
	player := createPlayer("name")

	if player.name != "name" {
		t.Errorf("Expected player name to be 'name' but got '%s' instead", player.name)
	}

	if player.score != 0 {
		t.Errorf("Expected player score to be 0 but got %d instead", player.score)
	}
}

// Test to see if resetting a players score works
func TestResetScore(t *testing.T) {
	player := createPlayer("name")

	player.resetScore()

	if player.score != 0 {
		t.Errorf("Expected player score to be 0 but got %d instead", player.score)
	}
}

// Test to see if increasing a players score works
func TestIncreaseScore(t *testing.T) {
	player := createPlayer("name")

	player.increaseScore(10)

	if player.score != 10 {
		t.Errorf("Expected score to be 10 but got %d instead", player.score)
	}
}
