package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const startCave = "start"
const endCave = "end"

type mapConnections map[string][]string

type path struct {
	currentCave string
	visited     map[string]int
}

func readConnections(producer *bufio.Scanner) mapConnections {

	connections := make(mapConnections)

	for producer.Scan() {
		lineInput := strings.Split(producer.Text(), "-")
		connections[lineInput[0]] = append(connections[lineInput[0]], lineInput[1])
		connections[lineInput[1]] = append(connections[lineInput[1]], lineInput[0])
	}

	return connections
}

func isBigCave(cave string) bool {
	return strings.ToUpper(cave) == cave
}

func buildVisited(source map[string]int, cave string) map[string]int {
	result := make(map[string]int, len(source)+1)

	for c, v := range source {
		result[c] = v
	}

	if !isBigCave(cave) {
		result[cave] += 1
	}

	return result
}

func doubleDip(source map[string]int, cave string) bool {
	if source[cave] == 0 {
		return false
	}

	if source[cave] == 1 {
		for _, v := range source {
			if v > 1 {
				return true
			}
		}

		return false
	}

	return true
}

func walkMap(mConns mapConnections) int {
	currentPaths := []path{{
		visited:     map[string]int{startCave: 1},
		currentCave: startCave,
	}}
	numPaths := 0

	for len(currentPaths) > 0 {
		var newPaths []path

		for _, curPath := range currentPaths {

			for _, possiblePath := range mConns[curPath.currentCave] {
				if possiblePath == endCave {
					numPaths += 1
					continue
				}

				if possiblePath == startCave {
					continue
				}

				if !isBigCave(possiblePath) && doubleDip(curPath.visited, possiblePath) {
					continue
				}

				fork := path{
					currentCave: possiblePath,
					visited:     buildVisited(curPath.visited, possiblePath),
				}

				newPaths = append(newPaths, fork)
			}
		}
		currentPaths = newPaths
	}

	return numPaths
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	conns := readConnections(scanner)
	return walkMap(conns)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
