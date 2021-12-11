package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const steps = 100

type octoMap [][]int

func readMap(producer *bufio.Scanner) (result octoMap) {

	for producer.Scan() {
		inputLine := strings.Split(producer.Text(), "")

		var mapLine []int
		for _, c := range inputLine {
			depth, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
			}
			mapLine = append(mapLine, depth)
		}

		result = append(result, mapLine)
	}

	return
}

func flash(oMap *octoMap, row, column int) {
	(*oMap)[row][column] = -1

	// increase neighbors unless they are -1 (already flashed this step)
	possibleRows := []int{row}
	possibleColumns := []int{column}

	if row > 0 {
		possibleRows = append(possibleRows, row-1)
	}
	if row < len(*oMap)-1 {
		possibleRows = append(possibleRows, row+1)
	}

	if column > 0 {
		possibleColumns = append(possibleColumns, column-1)
	}
	if column < len((*oMap)[0])-1 {
		possibleColumns = append(possibleColumns, column+1)
	}

	for _, r := range possibleRows {
		for _, c := range possibleColumns {
			if (*oMap)[r][c] != -1 {
				(*oMap)[r][c] += 1
			}
		}
	}
}

func stepMap(oMap *octoMap) (result int) {

	// increase all cells by one
	for i := range *oMap {
		for j := range (*oMap)[i] {
			(*oMap)[i][j] += 1
		}
	}

OUTER:
	for {
		for i := range *oMap {
			for j := range (*oMap)[i] {
				if (*oMap)[i][j] > 9 {
					flash(oMap, i, j)
					result += 1
					continue OUTER
				}
			}
		}
		break
	}

	// set flashed to 0
	for i := range *oMap {
		for j := range (*oMap)[i] {
			if (*oMap)[i][j] == -1 {
				(*oMap)[i][j] = 0
			}
		}
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	oMap := readMap(scanner)

	flashes := 0

	for i := 0; i < steps; i++ {
		flashes += stepMap(&oMap)
	}
	return flashes
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
