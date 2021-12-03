package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

type counters struct {
	ones, zeros int
}

func generateCounters(producer *bufio.Scanner) []counters {
	var store *[]counters
	for producer.Scan() {
		commandInput := producer.Text()

		for i := 0; i < len(commandInput); i++ {
			if store == nil {
				newStore := make([]counters, len(commandInput))
				store = &newStore
			}
			switch commandInput[i] {
			case '1':
				(*store)[i].ones += 1
			case '0':
				(*store)[i].zeros += 1
			default:
				log.Fatal("unknown value: ", commandInput[i])
			}
		}
	}

	return *store
}

func generateValues(store []counters) (gamma, epsilon int) {
	for i := 0; i < len(store); i++ {
		if store[i].ones > store[i].zeros {
			gamma += int(math.Pow(2, float64(len(store)-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(len(store)-i-1)))
		}
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	store := generateCounters(scanner)
	gamma, epsilon := generateValues(store)
	return gamma * epsilon
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Final position", result)
}
