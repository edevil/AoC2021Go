package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	hPos := 0
	vPos := 0
	aim := 0
	for scanner.Scan() {
		commandInput := scanner.Text()

		components := strings.Split(commandInput, " ")
		command := components[0]
		val, err := strconv.Atoi(components[1])
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case command == "forward":
			hPos += val
			vPos += val * aim
		case command == "up":
			aim -= val
		case command == "down":
			aim += val
		}
	}

	fmt.Println("Final position", hPos*vPos, " (", hPos, ",", vPos, ")")
}
