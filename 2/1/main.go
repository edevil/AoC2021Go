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
		case command == "up":
			vPos -= val
		case command == "down":
			vPos += val
		}

	}

	fmt.Println("Final position", hPos*vPos, " (", hPos, ",", vPos, ")")
}
