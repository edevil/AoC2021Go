package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const numSteps = 10

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

func doStep(s *state) {
	var newTemplate string
	for i := 0; i < len((*s).template)-1; i++ {
		fetcher := (*s).template[i : i+2]
		replacement := s.insertionRules[fetcher]
		newTemplate = fmt.Sprintf("%s%c%s", newTemplate, (*s).template[i], replacement)
	}
	s.template = fmt.Sprintf("%s%c", newTemplate, (*s).template[len((*s).template)-1])
}

func calculateResult(output string) int {
	occurrences := make(map[rune]int)

	for _, c := range output {
		occurrences[c] += 1
	}

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

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	s := readState(scanner)

	for i := 0; i < numSteps; i++ {
		doStep(&s)
	}

	return calculateResult(s.template)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
