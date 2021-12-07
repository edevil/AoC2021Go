package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type aquarium []int

var /* const */ separatorRE = regexp.MustCompile(`,|\s+`)

func readInitialState(producer *bufio.Scanner) aquarium {

	producer.Scan()
	return stringToNumbers(producer.Text())
}

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

func distanceToCost(distance int) (result int) {
	for i := 1; i <= distance; i++ {
		result += i
	}

	return
}

func calculateCost(aq aquarium, pos int) (result int) {

	for _, fishPos := range aq {
		distance := int(math.Abs(float64(pos) - float64(fishPos)))
		result += distanceToCost(distance)
	}

	return
}

func findMax(aq aquarium) (result int) {

	for _, fishPos := range aq {
		if fishPos > result {
			result = fishPos
		}
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	aq := readInitialState(scanner)

	maxPos := findMax(aq)
	bestCost := -1
	bestPos := -1
	for i := 0; i <= maxPos; i++ {
		if bestCost == -1 || calculateCost(aq, i) < bestCost {
			bestCost = calculateCost(aq, i)
			bestPos = i
		}
	}

	return calculateCost(aq, bestPos)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
