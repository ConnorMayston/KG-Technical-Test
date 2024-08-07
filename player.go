package main

type player struct {
	name  string
	score int
}

// Constructor for player
func createPlayer(name string) player {
	return player{
		name:  name,
		score: 0,
	}
}

// Resets player score
func (p *player) resetScore() {
	p.score = 0
}

// Increases player score
func (p *player) increaseScore(s int) {
	p.score += s
}
