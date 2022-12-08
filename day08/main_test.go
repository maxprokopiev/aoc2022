package main

import (
	"log"
	"os"
	"testing"
)

var testInput = `30373
25512
65332
33549
35390`

func TestSolve2TestInput(t *testing.T) {
	result := Solve2(testInput)
	if result != 8 {
		t.Errorf("got %v, want %v", result, 8)
	}
}

func TestSolve2Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve2(string(input))
	if result != 410400 {
		t.Errorf("got %v, want %v", result, 410400)
	}
}

func TestSolve1TestInput(t *testing.T) {
	result := Solve1(testInput)
	if result != 21 {
		t.Errorf("got %v, want %v", result, 21)
	}
}

func TestSolve1Input(t *testing.T) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := Solve1(string(input))
	if result != 1688 {
		t.Errorf("got %v, want %v", result, 1688)
	}
}
