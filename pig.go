package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pig struct {
	players         []player
	playerTurnIndex int
	dice            dice
}

func createGame(players []player) pig {
	return pig{
		players:         players,
		dice:            newDice(),
		playerTurnIndex: 0,
	}
}

func (p *pig) play() {

	for {
		endTurn := false
		sum := 0

		for !endTurn {
			fmt.Printf("\n%s, it is your turn. Do you want to a) roll or b) skip \nYour current score is: %d\n", p.getCurrentPlayer().name, p.getCurrentPlayer().score+sum)

			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')

			answer = strings.TrimRight(answer, "\n")

			for !(answer == "a" || answer == "b") {
				fmt.Println("Please enter either 'a' or 'b'")
				answer, _ = reader.ReadString('\n')
			}

			if answer == "a" {
				onesRolled := 0
				for i := 0; i < 2; i++ {
					roll := p.dice.roll()
					if roll == 1 {
						onesRolled++
					}
					sum += roll
				}
				endTurn = p.optionA(&sum, onesRolled)
			} else if answer == "b" {
				p.optionB(sum)
				endTurn = true
			}
		}
	}
}

func (p *pig) optionA(sum *int, onesRolled int) bool {

	if p.decideEndTurn(onesRolled, *sum) {
		return true
	}

	fmt.Printf("%s's score is now %d\n\n", p.getCurrentPlayer().name, p.getCurrentPlayer().score+(*sum))
	return false
}

func (p *pig) optionB(sum int) {
	p.getCurrentPlayer().increaseScore(sum)
	p.nextPlayer()
}

func (p *pig) decideEndTurn(onesRolled int, sum int) bool {
	if onesRolled == 1 {
		fmt.Printf("\nBecause you rolled a ONE. Your score has been reset to %d.\n", p.getCurrentPlayer().score)
		p.nextPlayer()
		return true
	}

	if onesRolled == 2 {
		fmt.Println("\nBecause you rolled snakes eyes (Two one's). Your score has been reset to 0!.")
		p.getCurrentPlayer().resetScore()
		p.nextPlayer()
		return true
	}

	//Can likely be taken a level higher and then sum doesn't need to be passed down and easier testing?
	if (sum + p.getCurrentPlayer().score) > 100 {
		p.win()
	}
	return false
}

func (p *pig) nextPlayer() {
	if p.playerTurnIndex == len(p.players)-1 {
		p.playerTurnIndex = 0
	} else {
		p.playerTurnIndex++
	}
}

func (p *pig) win() {
	fmt.Printf("\nCongratulations %s! You have won!", p.getCurrentPlayer().name)
	os.Exit(0)
}

func (p *pig) getCurrentPlayer() *player {
	return &p.players[p.playerTurnIndex]
}
