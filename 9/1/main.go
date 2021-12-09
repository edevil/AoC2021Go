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

type heightmap [][]int

func readMap(producer *bufio.Scanner) (result heightmap) {

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

func findLowPoints(hMap heightmap) (result []int) {
	for i := range hMap {
		for j := range hMap[i] {
			// check left
			if i != 0 && hMap[i-1][j] <= hMap[i][j] {
				continue
			}

			// check right
			if i != len(hMap)-1 && hMap[i+1][j] <= hMap[i][j] {
				continue
			}

			// check up
			if j != 0 && hMap[i][j-1] <= hMap[i][j] {
				continue
			}

			// check down
			if j != len(hMap[i])-1 && hMap[i][j+1] <= hMap[i][j] {
				continue
			}

			// low point
			result = append(result, hMap[i][j])
		}
	}

	return
}

func sumRiskLevels(lowPoints []int) (result int) {

	for _, val := range lowPoints {
		result += val + 1
	}
	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	hMap := readMap(scanner)
	lowPoints := findLowPoints(hMap)
	return sumRiskLevels(lowPoints)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
