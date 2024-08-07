package main

import (
	"testing"
)

func TestCreateGame(t *testing.T) {
	p := createGame([]player{
		{
			name:  "name",
			score: 0,
		},
		{
			name:  "name2",
			score: 0,
		},
		{
			name:  "name3",
			score: 0,
		},
	})

	if len(p.dice) != 6 {
		t.Errorf("Expected length of dice to be 6 instead got %d", len(p.dice))
	}

	if len(p.players) != 3 {
		t.Errorf("Expected length of players to be 3 instead got %d", len(p.players))
	}

	if p.players[0].name != "name" {
		t.Errorf("Expected player 1 name to be 'name' instead got %s", p.players[0].name)
	}

	if p.players[1].name != "name2" {
		t.Errorf("Expected player 2 name to be 'name2' instead got %s", p.players[1].name)
	}

	if p.players[2].name != "name3" {
		t.Errorf("Expected player 3 name to be 'name3' instead got %s", p.players[2].name)
	}

	if p.playerTurnIndex != 0 {
		t.Errorf("Expected the turn index to be 0 instead got %d", p.playerTurnIndex)
	}
}

func TestOptionA(t *testing.T) {
	p := createGame([]player{
		createPlayer("name1"),
		createPlayer("name2"),
	})

	sum := 0
	onesRolled := 1

	res := p.optionA(&sum, onesRolled)

	if res != true {
		t.Errorf("Expected res to be true instead got %t", res)
	}

	if sum != 0 {
		t.Errorf("Expected sum to be 0 instead got %d", sum)
	}

	onesRolled = 0
	sum = 20

	res = p.optionA(&sum, onesRolled)

	if res != false {
		t.Errorf("Expected res to be false instead got %t", res)
	}

	if sum != 20 {
		t.Errorf("Expected sum to be 20 instead got %d", sum)
	}
}

func TestOptionB(t *testing.T) {
	p := createGame([]player{
		createPlayer("name1"),
		createPlayer("name2"),
	})

	p.optionB(30)

	if p.players[0].score != 30 {
		t.Errorf("Expected player 0's score to be 30 instead got %d", p.players[0].score)
	}

	if p.playerTurnIndex != 1 {
		t.Errorf("Expected current player turn index to be 1 instead got %d", p.playerTurnIndex)
	}

	p.optionB(10)

	if p.players[1].score != 10 {
		t.Errorf("Expected player 1's score to be 10 instead got %d", p.players[1].score)
	}

	if p.playerTurnIndex != 0 {
		t.Errorf("Expected current player turn index to be 0 instead got %d", p.playerTurnIndex)
	}
}

func TestDecideEndTurn(t *testing.T) {
	p := createGame([]player{
		createPlayer("name1"),
		createPlayer("name2"),
	})

	p.players[0].increaseScore(20)

	endTurn := p.decideEndTurn(0, 10)

	if endTurn != false {
		t.Errorf("Expected endTurn to be false instead got %t", endTurn)
	}

	if p.players[0].score != 20 {
		t.Errorf("Expected player 0 score to be 20 instead got %d", p.players[0].score)
	}

	endTurn = p.decideEndTurn(1, 0)

	if endTurn != true {
		t.Errorf("Expected endTurn to be true instead got %t", endTurn)
	}

	if p.players[0].score != 20 {
		t.Errorf("Expected player 0 score to be 20 instead got %d", p.players[0].score)
	}

	p.nextPlayer()

	endTurn = p.decideEndTurn(2, 0)

	if endTurn != true {
		t.Errorf("Expected endTurn to be true instead got %t", endTurn)
	}

	if p.players[0].score != 0 {
		t.Errorf("Expected player 0 score to be 0 instead got %d", p.players[0].score)
	}

}

func TestNextPlayer(t *testing.T) {
	p := createGame([]player{
		createPlayer("name1"),
		createPlayer("name2"),
	})

	p.nextPlayer()

	if p.getCurrentPlayer().name != "name2" {
		t.Errorf("Expected current player's name to be 'name2' instead got %s", p.getCurrentPlayer().name)
	}

	p.nextPlayer()

	if p.getCurrentPlayer().name != "name1" {
		t.Errorf("Expected current player's name to be 'name1' instead got %s", p.getCurrentPlayer().name)
	}

}
