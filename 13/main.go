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

type dot struct {
	x, y int
}

type instruction struct {
	horizontal bool
	line       int
}

func readInstructions(producer *bufio.Scanner) (dots []dot, instructions []instruction) {

	for producer.Scan() {
		lineInput := strings.Split(producer.Text(), ",")
		if len(lineInput) == 1 {
			break
		}

		x, err := strconv.Atoi(lineInput[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(lineInput[1])
		if err != nil {
			log.Fatal(err)
		}

		inst := dot{
			x: x,
			y: y,
		}

		dots = append(dots, inst)
	}

	for producer.Scan() {
		var orientation string
		var line int
		_, err := fmt.Sscanf(producer.Text(), "fold along %1s=%d", &orientation, &line)
		if err != nil {
			log.Fatal(err)
		}

		i := instruction{
			line: line,
		}
		if orientation == "y" {
			i.horizontal = true
		}

		instructions = append(instructions, i)
	}

	return
}

func foldDots(dots []dot, inst instruction) []dot {
	dotMap := make(map[dot]bool)
	for _, d := range dots {
		if inst.horizontal && d.y > inst.line {
			d.y = d.y - 2*(d.y-inst.line)
		} else if !inst.horizontal && d.x > inst.line {
			d.x = d.x - 2*(d.x-inst.line)
		}
		dotMap[d] = true
	}

	var result []dot
	for k := range dotMap {
		result = append(result, k)
	}

	return result
}

func printDots(dots []dot) {
	var xMax, yMax int

	for _, d := range dots {
		if d.x > xMax {
			xMax = d.x
		}
		if d.y > yMax {
			yMax = d.y
		}
	}

	canvas := make([][]rune, yMax+1)
	for r := range canvas {
		canvas[r] = make([]rune, xMax+1)

		for c := range canvas[r] {
			canvas[r][c] = '.'
		}
	}

	// fill dots
	for _, d := range dots {
		canvas[d.y][d.x] = '#'
	}

	// print final
	for r := range canvas {
		fmt.Println(string(canvas[r]))
	}

}

func doIt(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	dots, insts := readInstructions(scanner)

	for _, i := range insts {
		dots = foldDots(dots, i)
	}

	printDots(dots)

	return len(dots)
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
