package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("WELCOME to Pig! Please enter 2 names in the following prompts to get started!")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("First Name: ")
	name, _ := reader.ReadString('\n')
	player1 := createPlayer(strings.TrimRight(name, "\n"))

	fmt.Print("Second Name: ")
	name, _ = reader.ReadString('\n')
	player2 := createPlayer(strings.TrimRight(name, "\n"))

	pig := createGame([]player{player1, player2})

	pig.play()
}
