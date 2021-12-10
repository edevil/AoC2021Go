package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
				return 3
			}
			state = state[:len(state)-1]
		case ']':
			if state[len(state)-1] != '[' {
				return 57
			}
			state = state[:len(state)-1]
		case '}':
			if state[len(state)-1] != '{' {
				return 1197
			}
			state = state[:len(state)-1]
		case '>':
			if state[len(state)-1] != '<' {
				return 25137
			}
			state = state[:len(state)-1]
		}
	}
	return 0
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	nSS := readNavigation(scanner)

	var result int
	for _, line := range nSS {
		result += calculateError(line)
	}

	return result
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
