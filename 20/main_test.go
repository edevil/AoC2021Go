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
	if result != 3351 {
		t.Errorf("doIt = %d; want 3351", result)
	}
}
