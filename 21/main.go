package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const dieSides = 3
const numSpaces = 10
const winningPoints = 21

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

	currentGames := map[gameState]int{game: 1}
	p1Wins := 0
	p2Wins := 0

	var possibleDieVals []int
	for i := 1; i <= dieSides; i++ {
		for j := 1; j <= dieSides; j++ {
			for k := 1; k <= dieSides; k++ {
				possibleDieVals = append(possibleDieVals, i+j+k)
			}
		}
	}

	for len(currentGames) > 0 {

		newGames := make(map[gameState]int)

		for game, gameCount := range currentGames {
			for _, dieValue := range possibleDieVals {
				nGame := game

				nGame.p1Start = (nGame.p1Start + dieValue) % numSpaces
				nGame.p1points += nGame.p1Start + 1

				if nGame.p1points >= winningPoints {
					p1Wins += gameCount
					continue
				}

				for _, die2Value := range possibleDieVals {
					// I know, I know. Should have a player array
					nnGame := nGame

					nnGame.p2Start = (nnGame.p2Start + die2Value) % numSpaces
					nnGame.p2points += nnGame.p2Start + 1

					if nnGame.p2points >= winningPoints {
						p2Wins += gameCount
						continue
					}

					newGames[nnGame] += gameCount
				}

			}

		}

		currentGames = newGames

	}

	if p1Wins > p2Wins {
		return p1Wins
	}

	return p2Wins
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
