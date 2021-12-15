package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

const mult = 5

type riskmap [][]int

func sumArray(source []int, inc int) (result []int) {
	for i := range source {
		newValue := source[i] + inc
		if newValue > 9 {
			newValue = newValue - 9
		}
		result = append(result, newValue)
	}

	return
}

func expandMap(rmap riskmap) riskmap {

	// expand horizontally
	for y := range rmap {
		var newRow []int
		for x := 0; x < mult; x++ {
			newRow = append(newRow, sumArray(rmap[y], x)...)
		}

		rmap[y] = newRow
	}

	// expand vertically
	numRows := len(rmap)
	for x := 1; x < mult; x++ {
		for r := 0; r < numRows; r++ {
			rmap = append(rmap, sumArray(rmap[r], x))
		}
	}

	return rmap
}

func readMap(producer *bufio.Scanner) (result riskmap) {

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

func vertIndex(rmap riskmap, x, y int) int {
	return x + y*len(rmap[y])
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	rmap := readMap(scanner)
	rmap = expandMap(rmap)
	graph := dijkstra.NewGraph()

	// add vertexes
	for y := range rmap {
		for x := range rmap[y] {
			graph.AddVertex(vertIndex(rmap, x, y))
		}
	}

	// add arcs
	for y := range rmap {
		for x := range rmap[y] {

			cVertex := vertIndex(rmap, x, y)
			// try north
			if y > 0 {
				weight := rmap[y-1][x]
				err := graph.AddArc(cVertex, vertIndex(rmap, x, y-1), int64(weight))
				if err != nil {
					log.Fatal(err)
				}
			}

			// try south
			if y < len(rmap)-1 {
				weight := rmap[y+1][x]
				err := graph.AddArc(cVertex, vertIndex(rmap, x, y+1), int64(weight))
				if err != nil {
					log.Fatal(err)
				}
			}

			// try east
			if x > 0 {
				weight := rmap[y][x-1]
				err := graph.AddArc(cVertex, vertIndex(rmap, x-1, y), int64(weight))
				if err != nil {
					log.Fatal(err)
				}
			}

			// try west
			if x < len(rmap[y])-1 {
				weight := rmap[y][x+1]
				err := graph.AddArc(cVertex, vertIndex(rmap, x+1, y), int64(weight))
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}

	best, err := graph.Shortest(0, vertIndex(rmap, len(rmap[0])-1, len(rmap)-1))
	if err != nil {
		log.Fatal(err)
	}

	return int(best.Distance)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
