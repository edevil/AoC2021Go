package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type heightmap [][]int

type point struct {
	x, y int
}

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

func findLowPoints(hMap heightmap) (result []point) {
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
			result = append(result, point{x: i, y: j})
		}
	}

	return
}

func sizeBasin(hMap heightmap, lowPoint point) int {
	result := 0

	pointsToCheck := map[point]bool{lowPoint: true}
	pointsChecked := make(map[point]bool)

	for len(pointsToCheck) > 0 {
		for p := range pointsToCheck {
			pointsChecked[p] = true
			delete(pointsToCheck, p)

			if hMap[p.x][p.y] == 9 {
				continue
			}

			result += 1

			// look left
			if p.x > 0 {
				newPoint := point{x: p.x - 1, y: p.y}
				if _, ok := pointsChecked[newPoint]; !ok {
					pointsToCheck[newPoint] = true
				}
			}

			// look right
			if p.x < len(hMap)-1 {
				newPoint := point{x: p.x + 1, y: p.y}
				if _, ok := pointsChecked[newPoint]; !ok {
					pointsToCheck[newPoint] = true
				}
			}

			// look up
			if p.y > 0 {
				newPoint := point{x: p.x, y: p.y - 1}
				if _, ok := pointsChecked[newPoint]; !ok {
					pointsToCheck[newPoint] = true
				}
			}

			// look down
			if p.y < len(hMap[p.x])-1 {
				newPoint := point{x: p.x, y: p.y + 1}
				if _, ok := pointsChecked[newPoint]; !ok {
					pointsToCheck[newPoint] = true
				}
			}
		}
	}

	return result
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	hMap := readMap(scanner)
	lowPoints := findLowPoints(hMap)

	var sizes []int
	for _, point := range lowPoints {
		sizes = append(sizes, sizeBasin(hMap, point))
	}

	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
