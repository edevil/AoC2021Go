package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type entry struct {
	patterns [10]string
	oValues  [4]string
}

func readInitialState(producer *bufio.Scanner) (result []entry) {

	for producer.Scan() {
		commandInput := strings.Split(producer.Text(), " | ")
		newEntry := entry{}
		copy(newEntry.patterns[:], strings.Split(commandInput[0], " "))
		copy(newEntry.oValues[:], strings.Split(commandInput[1], " "))
		result = append(result, newEntry)
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	entries := readInitialState(scanner)

	result := 0
	for _, e := range entries {
		for _, value := range e.oValues {
			switch len(value) {
			case 2, 3, 4, 7:
				result += 1
			}
		}
	}
	return result
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
