package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(l *list.List) int {
	sumVal := 0
	for e := l.Front(); e != nil; e = e.Next() {
		sumVal += e.Value.(int)
	}

	return sumVal
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	increases := 0
	prevWindow := list.New()
	currentWindow := list.New()
	for scanner.Scan() {
		iVal, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if currentWindow.Len() > 0 {
			prevWindow.PushBack(currentWindow.Back().Value)
		}

		currentWindow.PushBack(iVal)
		if currentWindow.Len() == 1 {
			continue
		}
		if currentWindow.Len() == 4 {
			currentWindow.Remove(currentWindow.Front())
		}

		if prevWindow.Len() == 4 {
			prevWindow.Remove(prevWindow.Front())
		}

		if prevWindow.Len() != 3 || currentWindow.Len() != 3 {
			continue
		}

		if sum(currentWindow) > sum(prevWindow) {
			increases += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of increases: ", increases)
}
