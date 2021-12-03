package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type counters struct {
	ones, zeros int
}

func readAllValues(producer *bufio.Scanner) []string {
	allValues := []string{}
	for producer.Scan() {
		commandInput := producer.Text()
		allValues = append(allValues, commandInput)
	}

	return allValues
}

func generateCounters(allValues []string) []counters {
	var store *[]counters
	for _, commandInput := range allValues {

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

func generateValues(allValues []string) (oxygen, co2 int) {
	optionsO2 := append([]string{}, allValues...)
	optionsCO2 := append([]string{}, allValues...)

	o2Store := generateCounters(optionsO2)
	co2Store := generateCounters(optionsCO2)
	numBits := len(o2Store)

	for i := 0; i < numBits; i++ {

		if len(optionsO2) != 1 {
			if o2Store[i].ones >= o2Store[i].zeros {
				optionsO2 = filterValues(optionsO2, i, '1')
			} else {
				optionsO2 = filterValues(optionsO2, i, '0')
			}

			if len(optionsO2) == 1 {
				oVal, err := strconv.ParseInt(optionsO2[0], 2, 32)
				if err != nil {
					log.Fatal(err)
				}
				oxygen = int(oVal)
			} else {
				o2Store = generateCounters(optionsO2)
			}
		}

		if len(optionsCO2) != 1 {
			if co2Store[i].ones >= co2Store[i].zeros {
				optionsCO2 = filterValues(optionsCO2, i, '0')
			} else {
				optionsCO2 = filterValues(optionsCO2, i, '1')
			}

			if len(optionsCO2) == 1 {
				cVal, err := strconv.ParseInt(optionsCO2[0], 2, 32)
				if err != nil {
					log.Fatal(err)
				}
				co2 = int(cVal)
			} else {
				co2Store = generateCounters(optionsCO2)
			}
		}

		if len(optionsO2) <= 1 && len(optionsCO2) <= 1 {
			break
		}

	}

	return oxygen, co2
}

func filterValues(possibleVals []string, idx int, filter byte) (result []string) {
	for _, val := range possibleVals {
		if val[idx] == filter {
			result = append(result, val)
		}
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	allValues := readAllValues(scanner)
	oxygen, co2 := generateValues(allValues)
	return oxygen * co2
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
