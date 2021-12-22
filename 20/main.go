package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const growFactor = 6
const idxFactor = growFactor / 2

type algorithm []bool
type canvas [][]int

func readInstructions(producer *bufio.Scanner) (alg algorithm, image canvas) {

	// read algorithm
	producer.Scan()
	instructions := producer.Text()
	alg = make([]bool, len(instructions))
	for i, val := range instructions {
		alg[i] = val == '#'
	}

	// read empty line
	producer.Scan()

	// read starting image
	for producer.Scan() {
		iRow := producer.Text()
		rowValues := make([]int, len(iRow))
		for i, val := range iRow {
			if val == '#' {
				rowValues[i] = 1
			}
		}
		image = append(image, rowValues)
	}

	return
}

func growImage(image canvas) (result canvas) {
	newHeight := len(image) + growFactor
	newWidth := len(image[0]) + growFactor
	result = make([][]int, newHeight)

	for i := 0; i < len(result); i++ {
		result[i] = make([]int, newWidth)
	}

	return
}

func getPixelNumber(image canvas, row, column, outside int) (result int) {

	for i := 0; i < idxFactor; i++ {
		for j := 0; j < idxFactor; j++ {
			result <<= 1
			r := row - 1 + i
			c := column - 1 + j
			if r < 0 || r >= len(image) || c < 0 || c >= len(image[r]) {
				result += outside
				continue
			}
			result += image[r][c]
		}
	}

	return
}

func stepImage(alg algorithm, image canvas, stepnumber int) (result canvas) {
	result = growImage(image)
	outside := 0
	if stepnumber%2 == 0 {
		outside = 1
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			pxNum := getPixelNumber(image, i-idxFactor, j-idxFactor, outside)
			if alg[pxNum] {
				result[i][j] = 1
			}
		}
	}

	return
}

func countLights(image canvas) (result int) {
	for i := 0; i < len(image); i++ {
		row := image[i]
		for j := 0; j < len(row); j++ {
			if row[j] == 1 {
				result += 1
			}
		}
	}

	return
}

func printImage(image canvas) {
	for r := range image {
		for c := range image[r] {
			if image[r][c] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	alg, image := readInstructions(scanner)
	printImage(image)

	for i := 1; i < 51; i++ {
		outside := 1
		if alg[0] {
			outside = i
		}
		image = stepImage(alg, image, outside)
		printImage(image)
	}

	return countLights(image)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
