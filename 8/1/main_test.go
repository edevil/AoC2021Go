package main

import (
	"log"
	"os"
	"testing"
)

func Test_doIt(t *testing.T) {
	inputFile, err := os.Open("input_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 0 {
		t.Errorf("doIt = %d; want 0", result)
	}
}

func Test2_doIt(t *testing.T) {
	inputFile, err := os.Open("input2_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 26 {
		t.Errorf("doIt = %d; want 26", result)
	}
}
