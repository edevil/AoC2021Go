package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	increases := 0
	prevValue := -1
	for scanner.Scan() {
		iVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if prevValue != -1 && prevValue < iVal {
			increases += 1
		}

		prevValue = iVal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of increases: ", increases)
}
