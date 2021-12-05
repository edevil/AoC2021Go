package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type line struct {
	x1, y1, x2, y2 int
}

type gameInput []line

type gameBoard [][]int

func readHydroVents(producer *bufio.Scanner) (gInput gameInput, maxX, maxY int) {

	for producer.Scan() {
		var x1, y1, x2, y2 int
		lineInput := producer.Text()
		_, err := fmt.Sscanf(lineInput, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}
		gInput = append(gInput, line{x1: x1, y1: y1, x2: x2, y2: y2})
		if x1 > maxX {
			maxX = x1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y1 > maxY {
			maxY = y1
		}
		if y2 > maxY {
			maxY = y2
		}
	}

	return
}

func drawBoard(gInput gameInput, maxX, maxY int) gameBoard {
	board := make([][]int, maxY+1)
	for i := range board {
		board[i] = make([]int, maxX+1)
	}

	for _, line := range gInput {
		var start, end int
		if line.x1 == line.x2 {
			// vertical line
			if line.y1 < line.y2 {
				start = line.y1
				end = line.y2
			} else {
				start = line.y2
				end = line.y1
			}

			for i := start; i <= end; i++ {
				board[i][line.x1] += 1
			}
		} else if line.y1 == line.y2 {
			// horizontal line
			if line.x1 < line.x2 {
				start = line.x1
				end = line.x2
			} else {
				start = line.x2
				end = line.x1
			}

			for i := start; i <= end; i++ {
				board[line.y1][i] += 1
			}

		} else {
			continue
		}
	}

	return board
}

func countBoard(board gameBoard, min int) (result int) {
	for _, line := range board {
		for _, val := range line {
			if val >= min {
				result += 1
			}
		}
	}

	return
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	gInput, maxX, maxY := readHydroVents(scanner)
	board := drawBoard(gInput, maxX, maxY)
	return countBoard(board, 2)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
