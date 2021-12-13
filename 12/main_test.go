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
	if result != 36 {
		t.Errorf("doIt = %d; want 36", result)
	}
}

func Test2_doIt(t *testing.T) {
	inputFile, err := os.Open("input2_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 103 {
		t.Errorf("doIt = %d; want 103", result)
	}
}

func Test3_doIt(t *testing.T) {
	inputFile, err := os.Open("input3_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 3509 {
		t.Errorf("doIt = %d; want 3509", result)
	}
}
