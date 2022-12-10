package main

import (
	"log"
	"os"
	"testing"
)

func TestSolve1TestInput(t *testing.T) {
	input, err := os.ReadFile("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 13140 {
		t.Errorf("got %v, want %v", result, 13140)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 11720 {
		t.Errorf("got %v, want %v", result, 11720)
	}
}
