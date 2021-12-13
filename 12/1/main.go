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
	visited     map[string]bool
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

func buildVisited(source map[string]bool, cave string) map[string]bool {
	result := make(map[string]bool, len(source)+1)

	for c := range source {
		result[c] = true
	}
	result[cave] = true

	return result
}

func walkMap(mConns mapConnections) int {
	currentPaths := []path{{
		visited:     map[string]bool{startCave: true},
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

				if !isBigCave(possiblePath) && curPath.visited[possiblePath] {
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
