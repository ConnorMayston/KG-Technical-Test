package main

type player struct {
	name  string
	score int
}

func createPlayer(name string) player {
	return player{
		name:  name,
		score: 0,
	}
}

func (p *player) resetScore() {
	p.score = 0
}

func (p *player) increaseScore(s int) {
	p.score += s
}
