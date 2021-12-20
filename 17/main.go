package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type state struct {
	x1, x2, y1, y2 int
}

func readState(input string) state {
	var x1, x2, y1, y2 int

	_, err := fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	if err != nil {
		log.Fatal(err)
	}

	return state{
		x1,
		x2,
		y1,
		y2,
	}
}

func inTarget(cState state, vx, vy int) bool {
	var curX, curY int

	for {
		if curX >= cState.x1 && curX <= cState.x2 && curY >= cState.y1 && curY <= cState.y2 {
			return true
		}

		if vx != 0 {
			curX += vx

			if vx > 0 {
				vx -= 1
			} else {
				vx += 1
			}
		}

		curY += vy
		vy -= 1

		if curY < cState.y1 {
			return false
		}
	}
}

func doIt(input io.Reader) int {
	packetData, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}
	cState := readState(strings.TrimRight(string(packetData), "\r\n"))

	count := 0
	for i := -5000; i < 5000; i++ {
		for j := -5000; j < 5000; j++ {
			if inTarget(cState, j, i) {
				count += 1
			}

		}
	}

	return count
}

func main() {
	result := doIt(os.Stdin)

	fmt.Println("Result: ", result)
}
