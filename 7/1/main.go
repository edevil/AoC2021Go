package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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

func calculateCost(aq aquarium, pos int) (result int) {

	for _, fishPos := range aq {
		result += int(math.Abs(float64(pos) - float64(fishPos)))
	}

	return
}

func calculateMedian(aq aquarium) int {
	sort.Ints(aq)

	middle := len(aq) / 2
	if len(aq)%2 != 0 {
		return aq[middle]
	}

	return (aq[middle-1] + aq[middle]) / 2
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	aq := readInitialState(scanner)
	pos := calculateMedian(aq)
	return calculateCost(aq, pos)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
