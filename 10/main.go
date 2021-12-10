package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type navigation []string

func readNavigation(producer *bufio.Scanner) (result navigation) {
	for producer.Scan() {
		commandInput := producer.Text()
		result = append(result, commandInput)
	}

	return
}

func calculateError(line string) int {
	var state []rune

	for _, c := range line {
		switch c {
		case '(', '[', '{', '<':
			state = append(state, c)
		case ')':
			if state[len(state)-1] != '(' {
				return 0
			}
			state = state[:len(state)-1]
		case ']':
			if state[len(state)-1] != '[' {
				return 0
			}
			state = state[:len(state)-1]
		case '}':
			if state[len(state)-1] != '{' {
				return 0
			}
			state = state[:len(state)-1]
		case '>':
			if state[len(state)-1] != '<' {
				return 0
			}
			state = state[:len(state)-1]
		}
	}

	total := 0
	for i := len(state) - 1; i >= 0; i-- {
		total *= 5
		switch state[i] {
		case '(':
			total += 1
		case '[':
			total += 2
		case '{':
			total += 3
		case '<':
			total += 4
		}

	}
	return total
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	nSS := readNavigation(scanner)

	var results []int
	for _, line := range nSS {
		cError := calculateError(line)
		if cError != 0 {
			results = append(results, cError)
		}
	}

	sort.Ints(results)

	return results[len(results)/2]
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
