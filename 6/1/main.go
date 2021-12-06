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

const (
	numDays       = 80
	gestingPeriod = 2
	birthReset    = 6
)

type aquarium []int

var /* const */ separatorRE = regexp.MustCompile(`,|\s+`)

func readInitialState(producer *bufio.Scanner) aquarium {
	producer.Scan()
	return stringToNumbers(producer.Text())
}

func advanceDay(aq aquarium) aquarium {
	newFish := 0

	for i := range aq {
		aq[i] -= 1
		if aq[i] < 0 {
			aq[i] = birthReset
			newFish += 1
		}
	}

	for i := 0; i < newFish; i++ {
		aq = append(aq, birthReset+gestingPeriod)
	}

	return aq
}

func advanceDays(aq aquarium, numDays int) aquarium {
	for i := 0; i < numDays; i++ {
		aq = advanceDay(aq)
	}

	return aq
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

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	aq := readInitialState(scanner)
	aq = advanceDays(aq, numDays)
	return len(aq)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
