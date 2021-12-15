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

type riskmap [][]int

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
