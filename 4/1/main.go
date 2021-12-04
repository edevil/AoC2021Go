package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const CardSize = 5

type bPos struct {
	value int
	drawn bool
}

type card [5][5]bPos

var /* const */ separatorRE = regexp.MustCompile(`,|\s+`)

func stringToNumbers(input string) (draws []int) {
	for _, val := range separatorRE.Split(strings.TrimSpace(input), -1) {
		iVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		draws = append(draws, iVal)
	}

	return
}

func readCard(producer *bufio.Scanner) (card, error) {

	var result card
	for i := 0; i < CardSize; i++ {
		if !producer.Scan() {
			return result, fmt.Errorf("Could not read card")
		}

		for j, val := range stringToNumbers(producer.Text()) {
			result[i][j] = bPos{value: val}
		}
	}
	return result, nil
}

func readGameState(producer *bufio.Scanner) (draws []int, cards []card) {

	// read list of numbers to draw
	producer.Scan()
	draws = stringToNumbers(producer.Text())

	// read empty line
	producer.Scan()

	// read cards
	for {
		newCard, err := readCard(producer)
		if err != nil {
			log.Fatal(err)
		}

		cards = append(cards, newCard)

		if !producer.Scan() {
			break
		}
	}

	return
}

func drawCard(draw int, pCard *card) {
	for i := range pCard {
		for j := range pCard[i] {
			if pCard[i][j].value == draw {
				pCard[i][j].drawn = true
			}
		}
	}
}

func isWinner(pCard card) bool {
OUTER_LINE:
	for _, line := range pCard {
		for _, pos := range line {
			if !pos.drawn {
				continue OUTER_LINE
			}
		}
		// line
		return true
	}

OUTER_COLUMN:
	for i := 0; i < CardSize; i++ {
		for _, line := range pCard {
			if !line[i].drawn {
				continue OUTER_COLUMN
			}
		}
		// column
		return true
	}

	return false
}

func sumUnmarked(pCard card) (sum int) {
	for i := range pCard {
		for j := range pCard[i] {
			if !pCard[i][j].drawn {
				sum += pCard[i][j].value
			}
		}
	}

	return
}

func playGame(draws []int, cards []card) (int, card) {
	for _, draw := range draws {
		for i := range cards {
			drawCard(draw, &cards[i])
			if isWinner(cards[i]) {
				return draw, cards[i]
			}
		}
	}

	panic("No winner")
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	draws, cards := readGameState(scanner)
	lastDraw, winningCard := playGame(draws, cards)
	return lastDraw * sumUnmarked(winningCard)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
