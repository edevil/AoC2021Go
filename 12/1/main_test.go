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
	if result != 10 {
		t.Errorf("doIt = %d; want 10", result)
	}
}

func Test2_doIt(t *testing.T) {
	inputFile, err := os.Open("input2_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 19 {
		t.Errorf("doIt = %d; want 19", result)
	}
}

func Test3_doIt(t *testing.T) {
	inputFile, err := os.Open("input3_test")
	if err != nil {
		log.Fatal(err)
	}

	result := doIt(inputFile)
	if result != 226 {
		t.Errorf("doIt = %d; want 226", result)
	}
}
