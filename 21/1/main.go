package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const dieSides = 100
const dieTimes = 3
const numSpaces = 10
const winningPoints = 1000

type gameState struct {
	p1Start, p1points, p2Start, p2points int
}

func readInitialState(producer *bufio.Scanner) (game gameState) {

	producer.Scan()
	_, err := fmt.Sscanf(producer.Text(), "Player 1 starting position: %d", &game.p1Start)
	if err != nil {
		log.Fatal(err)
	}
	game.p1Start -= 1

	producer.Scan()
	_, err = fmt.Sscanf(producer.Text(), "Player 2 starting position: %d", &game.p2Start)
	if err != nil {
		log.Fatal(err)
	}
	game.p2Start -= 1

	return
}

func playGame(game gameState) int {

	turn := 0

	for {
		// p1
		var dieTotal int
		for i := 0; i < dieTimes; i++ {
			dieValue := (turn % dieSides) + 1
			dieTotal += dieValue
			turn += 1
		}
		game.p1Start = (game.p1Start + dieTotal) % numSpaces
		game.p1points += game.p1Start + 1

		if game.p1points >= winningPoints {
			return game.p2points * turn
		}

		// p2
		dieTotal = 0
		for i := 0; i < dieTimes; i++ {
			dieValue := (turn % dieSides) + 1
			dieTotal += dieValue
			turn += 1
		}
		game.p2Start = (game.p2Start + dieTotal) % numSpaces
		game.p2points += game.p2Start + 1

		if game.p2points >= winningPoints {
			return game.p1points * turn
		}
	}
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	game := readInitialState(scanner)
	return playGame(game)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
