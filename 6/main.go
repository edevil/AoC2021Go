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
	numDays       = 256
	gestingPeriod = 2
	birthReset    = 6
)

type aquarium []int

var /* const */ separatorRE = regexp.MustCompile(`,|\s+`)

func readInitialState(producer *bufio.Scanner) aquarium {
	aq := make([]int, gestingPeriod+birthReset+1)

	producer.Scan()
	fish := stringToNumbers(producer.Text())
	for _, fishAge := range fish {
		aq[fishAge] += 1
	}

	return aq
}

func advanceDay(aq aquarium) aquarium {
	newFish := aq[0]

	for i := 0; i < len(aq)-1; i++ {
		aq[i] = aq[i+1]
	}

	aq[birthReset] += newFish
	aq[birthReset+gestingPeriod] = newFish

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

func sumFish(aq aquarium) (result int) {

	for _, count := range aq {
		result += count
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	aq := readInitialState(scanner)
	aq = advanceDays(aq, numDays)
	return sumFish(aq)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
