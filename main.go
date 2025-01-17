package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Maybe modularized further as well? Not sure if there are many benefits to such though, possibly readability?
func main() {
	fmt.Println("WELCOME to Pig! Please enter the number of players!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Number of players: ")

	validNumber := false

	numPlayers := 0

	for !validNumber {
		numPlayersStr, _ := reader.ReadString('\n')

		var err error

		numPlayers, err = strconv.Atoi(strings.TrimRight(numPlayersStr, "\n"))

		if err != nil {
			fmt.Println("Please enter a valid integer")
		} else if numPlayers < 2 {
			fmt.Println("You have to have at least 2 players, please enter a number again")
		} else {
			validNumber = true
		}
	}

	players := []player{}

	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Please enter name %d: ", i+1)
		name, _ := reader.ReadString('\n')
		players = append(players, createPlayer(strings.TrimRight(name, "\n")))
	}

	pig := createGame(players)

	pig.play()
}
