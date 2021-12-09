package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type entry struct {
	patterns [10]string
	oValues  [4]string
}

type mapping map[string]int

func sortString(input string) string {
	runes := []rune(input)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func readInitialState(producer *bufio.Scanner) (result []entry) {

	for producer.Scan() {
		commandInput := strings.Split(producer.Text(), " | ")
		newEntry := entry{}
		copy(newEntry.patterns[:], strings.Split(commandInput[0], " "))
		copy(newEntry.oValues[:], strings.Split(commandInput[1], " "))

		for i := range newEntry.patterns {
			newEntry.patterns[i] = sortString(newEntry.patterns[i])
		}

		for i := range newEntry.oValues {
			newEntry.oValues[i] = sortString(newEntry.oValues[i])
		}

		result = append(result, newEntry)
	}

	return
}

func removeChars(source, charsToRemove string) string {
	filterF := func(r rune) rune {
		if strings.ContainsRune(charsToRemove, r) {
			return -1
		}
		return r
	}

	return strings.Map(filterF, source)
}

func containsAllChars(source, chars string) bool {
	for _, c := range chars {
		if !strings.ContainsRune(source, c) {
			return false
		}
	}

	return true
}

func generateMapping(patterns []string) mapping {
	result := make(map[string]int)
	inverse := make(map[int]string)

	var (
		cfPat string
		bdPat string
	)

	// find 1,4,7,8
	var patternsLeft []string
	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			result[pattern] = 1
			inverse[1] = pattern
		case 3:
			result[pattern] = 7
			inverse[7] = pattern
		case 4:
			result[pattern] = 4
			inverse[4] = pattern
		case 7:
			result[pattern] = 8
			inverse[8] = pattern
		default:
			patternsLeft = append(patternsLeft, pattern)
		}
	}

	cfPat = inverse[1]
	bdPat = removeChars(inverse[4], inverse[1])

	// find len 5, 6 patterns
	var pat5s, pat6s []string
	for _, pattern := range patternsLeft {
		switch len(pattern) {
		case 5:
			pat5s = append(pat5s, pattern)
		case 6:
			pat6s = append(pat6s, pattern)
		default:
			panic(fmt.Sprintf("Unexpected pattern length: %v", pattern))
		}
	}

	// solve len 5 patterns
	for _, pat := range pat5s {
		if containsAllChars(pat, bdPat) {
			result[pat] = 5
			inverse[5] = pat
		} else if containsAllChars(pat, cfPat) {
			result[pat] = 3
			inverse[3] = pat
		} else {
			result[pat] = 2
			inverse[2] = pat
		}
	}

	// solve len 6 patterns
	for _, pat := range pat6s {
		if containsAllChars(pat, inverse[4]) {
			result[pat] = 9
			inverse[9] = pat
		} else if containsAllChars(pat, inverse[1]) && !containsAllChars(pat, inverse[4]) {
			result[pat] = 0
			inverse[0] = pat
		} else if !containsAllChars(pat, inverse[1]) && !containsAllChars(pat, inverse[4]) {
			result[pat] = 6
			inverse[6] = pat
		}
	}
	return result
}

func decodeValue(eMapping mapping, values [4]string) (result int) {
	for i := 0; i < len(values); i++ {
		result = result*10 + eMapping[values[i]]
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	entries := readInitialState(scanner)

	result := 0
	for _, e := range entries {
		eMapping := generateMapping(e.patterns[:])
		result += decodeValue(eMapping, e.oValues)
	}
	return result
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
