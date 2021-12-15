package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const numSteps = 40

type state struct {
	template       string
	insertionRules map[string]string
}

func readState(producer *bufio.Scanner) state {
	var result state

	// read template
	producer.Scan()
	result.template = producer.Text()

	// read empty line
	producer.Scan()

	// read insertion rules
	result.insertionRules = make(map[string]string)
	for producer.Scan() {
		var match string
		var replace string
		_, err := fmt.Sscanf(producer.Text(), "%s -> %s", &match, &replace)
		if err != nil {
			log.Fatal(err)
		}

		result.insertionRules[match] = replace
	}

	return result
}

func calculateResult(occurrences map[rune]int) int {
	max := -1
	min := -1

	for _, v := range occurrences {
		if max == -1 || v > max {
			max = v
		}
		if min == -1 || v < min {
			min = v
		}
	}
	return max - min
}

func occurrencesFromPairs(pairs map[string]int, template string) map[rune]int {

	occurrences := make(map[rune]int)
	for pair, count := range pairs {
		occurrences[rune(pair[0])] += count
		occurrences[rune(pair[1])] += count
	}

	occurrences[rune(template[0])] += 1
	occurrences[rune(template[len(template)-1])] += 1

	for lrune := range occurrences {
		occurrences[lrune] /= 2
	}
	return occurrences
}

func obtainPairCount(template string) map[string]int {
	occurrences := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]
		occurrences[pair] += 1
	}

	return occurrences
}

func stepPairCount(pairs map[string]int, insertionRules map[string]string) map[string]int {
	result := make(map[string]int)
	for pair, count := range pairs {
		newLetter := insertionRules[pair]
		result[string([]byte{pair[0], newLetter[0]})] += count
		result[string([]byte{newLetter[0], pair[1]})] += count
	}

	return result
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	s := readState(scanner)

	pairCount := obtainPairCount(s.template)
	for i := 0; i < numSteps; i++ {
		pairCount = stepPairCount(pairCount, s.insertionRules)
	}

	occurrences := occurrencesFromPairs(pairCount, s.template)

	return calculateResult(occurrences)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
